package main

import (
	"github.com/kubeflow/katib/pkg/ui"
	"github.com/pressly/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	kuh := ui.NewKatibUIHandler()
	r.Route("/", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, req *http.Request) {
			http.Redirect(writer, req, "/katib", http.StatusMovedPermanently)
		})
		r.Handle("/katib/static/*", http.StripPrefix("/katib/static/", http.FileServer(http.Dir("/app/static"))))
		r.Route("/katib", func(r chi.Router) {
			r.Get("/", kuh.Index)
			r.Get("/studyjob", kuh.StudyJobGen)
			r.Post("/studyjob", kuh.CreateStudyJob)
			r.Get("/workertemplates", kuh.WorkerTemplate)
			r.Post("/workertemplates", kuh.UpdateWorkerTemplate)
			r.Get("/metricscollectortemplates", kuh.MetricsCollectorTemplate)
			r.Post("/metricscollectortemplates", kuh.UpdateMetricsCollectorTemplate)
			r.Route("/proxy", func(r chi.Router) {
				r.Get("/GetStudyJobList", kuh.GetStudyJobList)
			})
			r.Route("/{studyid}", func(r chi.Router) {
				r.Get("/", kuh.Study)
				r.Get("/csv", kuh.StudyInfoCsv)
				r.Route("/TrialID/{trialid}", func(r chi.Router) {
					r.Get("/", kuh.Trial)
				})
				r.Route("/WorkerID/{workerid}", func(r chi.Router) {
					r.Get("/", kuh.Worker)
					r.Get("/csv", kuh.WorkerInfoCsv)
				})
			})
		})
	})
	http.ListenAndServe(":80", r)
}
