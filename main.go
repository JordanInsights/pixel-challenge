package main

import (
	"pixel-challenge/orchestration"
	"time"
)

func main() {
	// go func() {
	// 	mux := http.NewServeMux()
	// 	mux.HandleFunc("/debug/pprof/", pprof.Index)
	// 	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// 	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// 	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// 	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// 	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("hello"))
	// 	})

	// 	http.ListenAndServe(":6060", mux)
	// }()

	start := time.Now()

	filepaths := orchestration.GetFilepathsFromCommandLineArguments()
	orchestration.RunAnalysis(filepaths)

	elapsed := time.Since(start)
	orchestration.OutputSimilarities(elapsed, filepaths)

	// time.Sleep(time.Second * 60)
}
