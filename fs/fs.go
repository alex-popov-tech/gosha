package fs

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Report struct {
	ResultsDir string
}

type Run struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Suite struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Test struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Step struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Steps  []Step `json:"steps"`
}

type path struct{}

func (it path) runs(root string) string {
	return fmt.Sprintf("%s/%s", root, "runs")
}

func (it path) run(root, id string) string {
	return fmt.Sprintf("%s/%s/%s", it.runs(root), id, "data.json")
}

func (it path) suites(root, runId string) string {
	return fmt.Sprintf("%s/%s/%s", it.runs(root), runId, "suites")
}

func (it path) suite(root, runId, suiteId string) string {
	return fmt.Sprintf("%s/%s/%s", it.suites(root, runId), suiteId, "data.json")
}

func (it path) tests(root, runId, suiteId string) string {
	return fmt.Sprintf("%s/%s/%s", it.suites(root, runId), suiteId, "tests")
}

func (it path) test(root, runId, suiteId, testId string) string {
	return fmt.Sprintf("%s/%s/%s", it.tests(root, runId, suiteId), testId, "data.json")
}

func (it path) steps(root, runId, suiteId, testId string) string {
	return fmt.Sprintf("%s/%s/%s", it.tests(root, runId, suiteId), testId, "steps")
}

func (it path) step(root, runId, suiteId, testId, stepId string) string {
	return fmt.Sprintf("%s/%s/%s", it.steps(root, runId, suiteId, testId), stepId, "data.json")
}

var p path = path{}

func (it *Report) GetRuns() ([]Run, error) {
	path := p.runs(it.ResultsDir)
	folders, err := ioutil.ReadDir(path)
	if err != nil {
		return []Run{}, errors.WithMessage(err, fmt.Sprintf("GetRuns failed trying to read %s dir's content", path))
	}

	runs := []Run{}
	for _, folder := range folders {
		run, err := it.GetRun(folder.Name())
		if err != nil {
			return nil, errors.Wrap(err, "GetRuns failed because of:")
		}
		runs = append(runs, run)
	}
	return runs, nil
}

func (it *Report) GetRun(runId string) (Run, error) {
	run := Run{}
	path := p.run(it.ResultsDir, runId)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return run, errors.WithMessage(err, fmt.Sprintf("GetRun failed trying to read %s", path))
	}
	err = json.Unmarshal(file, &run)
	if err != nil {
		return run, errors.WithMessage(err, fmt.Sprintf("GetRun failed trynig to parse %s", path))
	}
	return run, nil
}

func (it *Report) GetSuites(runId string) ([]Suite, error) {
	path := p.suites(it.ResultsDir, runId)
	folders, err := ioutil.ReadDir(path)
	if err != nil {
		return []Suite{}, errors.WithMessage(err, fmt.Sprintf("GetRunSuites failed trynig to read %s's contents", path))
	}
	suites := []Suite{}
	for _, folder := range folders {
		suite, err := it.GetSuite(runId, folder.Name())
		if err != nil {
			return nil, errors.Wrap(err, "GetSuites failed because of:")
		}
		suites = append(suites, suite)
	}
	return suites, nil
}

func (it *Report) GetSuite(runId, suiteId string) (Suite, error) {
	path := p.suite(it.ResultsDir, runId, suiteId)
	file, err := ioutil.ReadFile(path)
	suite := Suite{}
	if err != nil {
		return suite, errors.WithMessage(err, fmt.Sprintf("GetSuite failed trying to read %s", path))
	}
	err = json.Unmarshal(file, &suite)
	if err != nil {
		return suite, errors.WithMessage(err, fmt.Sprintf("GetSuite failed trynig to parse %s", path))
	}
	return suite, nil
}

func (it *Report) GetTests(runId, suiteId string) ([]Test, error) {
	path := p.tests(it.ResultsDir, runId, suiteId)
	folders, err := ioutil.ReadDir(path)
	tests := []Test{}
	if err != nil {
		return tests, errors.WithMessage(err, fmt.Sprintf("GetTests failed trynig to read %s's contents", path))
	}
	for _, folder := range folders {
		test, err := it.GetTest(runId, suiteId, folder.Name())
		if err != nil {
			return nil, errors.Wrap(err, "GetTests failed because of:")
		}
		tests = append(tests, test)
	}
	return tests, nil
}

func (it *Report) GetTest(runId, suiteId, testId string) (Test, error) {
	path := p.test(it.ResultsDir, runId, suiteId, testId)
	file, err := ioutil.ReadFile(path)
	test := Test{}
	if err != nil {
		return test, errors.WithMessage(err, fmt.Sprintf("GetTest failed trying to read %s", path))
	}
	err = json.Unmarshal(file, &test)
	if err != nil {
		return test, errors.WithMessage(err, fmt.Sprintf("GetTest failed trynig to parse %s", path))
	}
	return test, nil
}

func (it *Report) GetSteps(runId, suiteId, testId string) ([]Step, error) {
	path := p.steps(it.ResultsDir, runId, suiteId, testId)
	folders, err := ioutil.ReadDir(path)
	steps := []Step{}
	if err != nil {
		return steps, errors.WithMessage(err, fmt.Sprintf("GetTests failed trynig to read %s's contents", path))
	}
	for _, folder := range folders {
		step, err := it.GetStep(runId, suiteId, testId, folder.Name())
		if err != nil {
			return nil, errors.Wrap(err, "GetTests failed because of:")
		}
		steps = append(steps, step)
	}
	return steps, nil
}

func (it *Report) GetStep(runId, suiteId, testId, stepId string) (Step, error) {
	path := p.step(it.ResultsDir, runId, suiteId, testId, stepId)
	file, err := ioutil.ReadFile(path)
	step := Step{}
	if err != nil {
		return step, errors.WithMessage(err, fmt.Sprintf("GetStep failed trying to read %s", path))
	}
	err = json.Unmarshal(file, &step)
	if err != nil {
		return step, errors.WithMessage(err, fmt.Sprintf("GetStep failed trynig to parse %s", path))
	}
	return step, nil
}
