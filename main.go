package main

import (
	"embed"
	rf "gosha/fs"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

const RELATIVE_RESULTS_DIR = "/data"

var reportFs rf.Report = rf.Report{ResultsDir: "." + RELATIVE_RESULTS_DIR}

var assetsFS fs.FS

//go:embed dist/assets/*
var assetsEmbedFs embed.FS

//go:embed dist/index.html
var index []byte

func init() {
	assetsFS, _ = fs.Sub(assetsEmbedFs, "dist/assets")
}

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	// frontend assets
	router.StaticFS("/assets", http.FS(assetsFS))

	// entrypoint
	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", index)
	})

	// runs
	router.GET("/runs", func(c *gin.Context) {
		runs, err := reportFs.GetRuns()
		if err != nil {
			c.JSON(503, gin.H{"error": err.Error()})
		}
		c.JSON(200, runs)
	})
	router.GET("/runs/:runId", func(c *gin.Context) {
		run, err := reportFs.GetRun(c.Param("runId"))
		if err != nil {
			c.JSON(503, gin.H{"error": err})
		}
		c.JSON(200, run)
	})
	// suites
	router.GET("/runs/:runId/suites", func(c *gin.Context) {
		suites, err := reportFs.GetSuites(c.Param("runId"))
		if err != nil {
			c.JSON(503, gin.H{"error": err.Error()})
		}
		c.JSON(200, suites)
	})
	router.GET("/runs/:runId/suites/:suiteId", func(c *gin.Context) {
		suite, err := reportFs.GetSuite(c.Param("runId"), c.Param("suiteId"))
		if err != nil {
			c.JSON(503, gin.H{"error": err})
		}
		c.JSON(200, suite)
	})
	// tests
	router.GET("/runs/:runId/suites/:suiteId/tests", func(c *gin.Context) {
		tests, err := reportFs.GetTests(c.Param("runId"), c.Param("suiteId"))
		if err != nil {
			c.JSON(503, gin.H{"error": err})
		}
		c.JSON(200, tests)
	})
	router.GET("/runs/:runId/suites/:suiteId/tests/:testId", func(c *gin.Context) {
		test, err := reportFs.GetTest(c.Param("runId"), c.Param("suiteId"), c.Param("testId"))
		if err != nil {
			c.JSON(503, gin.H{"error": err})
		}
		c.JSON(200, test)
	})
	// steps
	router.GET("/runs/:runId/suites/:suiteId/tests/:testId/steps", func(c *gin.Context) {
		steps, err := reportFs.GetSteps(c.Param("runId"), c.Param("suiteId"), c.Param("testId"))
		if err != nil {
			c.JSON(503, gin.H{"error": err})
		}
		c.JSON(200, steps)
	})
	router.GET("/runs/:runId/suites/:suiteId/tests/:testId/steps/:stepId", func(c *gin.Context) {
		step, err := reportFs.GetStep(c.Param("runId"), c.Param("suiteId"), c.Param("testId"), c.Param("stepId"))
		if err != nil {
			c.JSON(503, gin.H{"error": err})
		}
		c.JSON(200, step)
	})

	router.Run(":8080")
}

// const suitesElement = document.querySelector('article');
// fetch("/runs")
//   .then(it => it.json())
//   .then(async (runs) => {
//     console.log({ runs });
//     for (const { id: runId } of runs) {
//       const run = await fetch(`/runs/${runId}`).then(it => it.json());
//       console.log({ run });
//       const suites = await fetch(`/runs/${runId}/suites`).then(it => it.json());
//       console.log({ suites });
//       for (const { id: suiteId } of suites) {
//         const suite = await fetch(`/runs/${runId}/suites/${suiteId}`).then(it => it.json());
//         console.log({ suite });
//         const tests = await fetch(`/runs/${runId}/suites/${suiteId}/tests`).then(it => it.json());
//         console.log({ tests });
//         for (const { id: testId } of tests) {
//           const test = await fetch(`/runs/${runId}/suites/${suiteId}/tests/${testId}`).then(it => it.json());
//           console.log({ test });
//           const steps = await fetch(`/runs/${runId}/suites/${suiteId}/tests/${testId}/steps`).then(it => it.json());
//           console.log({ steps });
//             for (const { id: stepId } of steps) {
//               const step = await fetch(`/runs/${runId}/suites/${suiteId}/tests/${testId}/steps/${stepId}`).then(it => it.json());
//               console.log({ step });
//             }
//         }
//       }
//     }
//   });
//
