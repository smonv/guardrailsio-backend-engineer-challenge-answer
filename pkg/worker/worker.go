package worker

import (
	"beca"
	"bufio"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Worker struct {
	Ctx           context.Context
	jobChan       chan int
	logger        echo.Logger
	ResultService beca.ResultService
}

func New(ctx context.Context, jobChan chan int, logger echo.Logger, resultService beca.ResultService) *Worker {
	return &Worker{Ctx: ctx, jobChan: jobChan, logger: logger, ResultService: resultService}
}

func (w *Worker) Run() {
	for resultId := range w.jobChan {
		result, err := w.ResultService.Result(resultId)
		if result == nil {
			w.logger.Errorf("result with id %d not found\n", resultId)
			continue
		}
		if err != nil {
			w.logger.Error(err)
			continue
		}

		now := time.Now()
		result.Status = "In Process"
		result.ScanningAt = &now

		err = w.ResultService.UpdateResult(result)
		if err != nil {
			w.logger.Error(err)
			continue
		}

		tmpDir, err := ioutil.TempDir(os.TempDir(), result.RepositoryName)
		if err != nil {
			w.logger.Error(err)
			continue
		}

		cmd := exec.Command("git", "clone", result.RepositoryURL, tmpDir)

		err = cmd.Run()
		if err != nil {
			w.logger.Error(err)

			now = time.Now()
			result.FinishedAt = &now
			result.Status = "Failure"

			err = w.ResultService.UpdateResult(result)
			if err != nil {
				w.logger.Error(err)
			}

			continue
		}

		findings, err := w.scan(tmpDir)

		now = time.Now()
		result.FinishedAt = &now

		if err != nil {
			w.logger.Error(err)

			result.Status = "Failure"

			err = w.ResultService.UpdateResult(result)
			if err != nil {
				w.logger.Error(err)
			}

			continue
		}

		result.Status = "Success"
		result.Findings = findings

		err = w.ResultService.UpdateResult(result)
		if err != nil {
			w.logger.Error(err)
			continue
		}

		err = os.RemoveAll(tmpDir)
		if err != nil {
			w.logger.Fatal(err)
			continue
		}
	}
}

func (w *Worker) scan(dir string) ([]*beca.Finding, error) {
	paths, err := collectFilePath(dir)
	if err != nil {
		return nil, err
	}

	totalFindings := []*beca.Finding{}

	for _, path := range paths {
		findings, err := scanFile(path)
		if err != nil {
			return nil, err
		}

		totalFindings = append(totalFindings, findings...)
	}

	return totalFindings, nil
}

func collectFilePath(path string) ([]string, error) {
	paths := []string{}
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() && f.Name() == ".git" {
			return filepath.SkipDir
		}

		if !f.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			contentType, err := getFileContentType(file)
			if err != nil {
				return err
			}

			if strings.Contains(contentType, "text/plain") || strings.Contains(contentType, "application/octet-stream") {
				paths = append(paths, path)
			}

			err = file.Close()
			if err != nil {
				return err
			}
		}

		return nil
	})

	return paths, err
}

func scanFile(path string) ([]*beca.Finding, error) {
	findings := []*beca.Finding{}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		for _, word := range strings.Split(strings.TrimSpace(scanner.Text()), " ") {
			if strings.HasPrefix(word, "public_key") || strings.HasPrefix(word, "private_key") {
				finding := &beca.Finding{
					FindingType: "sast",
					RuleId:      "G402",
					Location: &beca.FindingLocation{
						Path: path,
						Positions: &beca.FindingLocationPositions{
							Begin: &beca.FindingLocationPositionsBegin{Line: line},
						},
					},
					Metadata: &beca.FindingMetadata{
						Description: "SECRET_KEY detected",
						Severity:    "HIGH",
					},
				}

				findings = append(findings, finding)
			}
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return findings, nil
}

func getFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
