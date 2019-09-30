package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type NzbgetStatus struct {
	Result  NzbgetStatusResult `json:"result"`
	Version string             `json:"version"`
}

type NzbgetServer struct {
	Active bool  `json:"Active"`
	ID     int64 `json:"ID"`
}

type NzbgetStatusResult struct {
	ArticleCacheHi      int64                `json:"ArticleCacheHi"`
	ArticleCacheLo      int64                `json:"ArticleCacheLo"`
	ArticleCacheMB      int64                `json:"ArticleCacheMB"`
	AverageDownloadRate int64                `json:"AverageDownloadRate"`
	DaySizeHi           int64                `json:"DaySizeHi"`
	DaySizeLo           int64                `json:"DaySizeLo"`
	DaySizeMB           int64                `json:"DaySizeMB"`
	Download2Paused     bool                 `json:"Download2Paused"`
	DownloadLimit       int64                `json:"DownloadLimit"`
	DownloadPaused      bool                 `json:"DownloadPaused"`
	DownloadRate        int64                `json:"DownloadRate"`
	DownloadTimeSec     int64                `json:"DownloadTimeSec"`
	DownloadedSizeHi    int64                `json:"DownloadedSizeHi"`
	DownloadedSizeLo    int64                `json:"DownloadedSizeLo"`
	DownloadedSizeMB    int64                `json:"DownloadedSizeMB"`
	FeedActive          bool                 `json:"FeedActive"`
	ForcedSizeHi        int64                `json:"ForcedSizeHi"`
	ForcedSizeLo        int64                `json:"ForcedSizeLo"`
	ForcedSizeMB        int64                `json:"ForcedSizeMB"`
	FreeDiskSpaceHi     int64                `json:"FreeDiskSpaceHi"`
	FreeDiskSpaceLo     int64                `json:"FreeDiskSpaceLo"`
	FreeDiskSpaceMB     int64                `json:"FreeDiskSpaceMB"`
	MonthSizeHi         int64                `json:"MonthSizeHi"`
	MonthSizeLo         int64                `json:"MonthSizeLo"`
	MonthSizeMB         int64                `json:"MonthSizeMB"`
	NewsServers         []NzbgetServer `json:"NewsServers"`
	ParJobCount         int64                `json:"ParJobCount"`
	PostJobCount        int64                `json:"PostJobCount"`
	PostPaused          bool                 `json:"PostPaused"`
	QueueScriptCount    int64                `json:"QueueScriptCount"`
	QuotaReached        bool                 `json:"QuotaReached"`
	RemainingSizeHi     int64                `json:"RemainingSizeHi"`
	RemainingSizeLo     int64                `json:"RemainingSizeLo"`
	RemainingSizeMB     int64                `json:"RemainingSizeMB"`
	ResumeTime          int64                `json:"ResumeTime"`
	ScanPaused          bool                 `json:"ScanPaused"`
	ServerPaused        bool                 `json:"ServerPaused"`
	ServerStandBy       bool                 `json:"ServerStandBy"`
	ServerTime          int64                `json:"ServerTime"`
	ThreadCount         int64                `json:"ThreadCount"`
	UpTimeSec           int64                `json:"UpTimeSec"`
	URLCount            int64                `json:"UrlCount"`
}

type NzbgetQueue struct {
	Result  []NzbgetQueueResult `json:"result"`
	Version string                   `json:"version"`
}

type NzbgetQueueResult struct {
	ActiveDownloads    int64                    `json:"ActiveDownloads"`
	Category           string                   `json:"Category"`
	CriticalHealth     int64                    `json:"CriticalHealth"`
	DeleteStatus       string                   `json:"DeleteStatus"`
	Deleted            bool                     `json:"Deleted"`
	DestDir            string                   `json:"DestDir"`
	DownloadTimeSec    int64                    `json:"DownloadTimeSec"`
	DownloadedSizeHi   int64                    `json:"DownloadedSizeHi"`
	DownloadedSizeLo   int64                    `json:"DownloadedSizeLo"`
	DownloadedSizeMB   int64                    `json:"DownloadedSizeMB"`
	DupeKey            string                   `json:"DupeKey"`
	DupeMode           string                   `json:"DupeMode"`
	DupeScore          int64                    `json:"DupeScore"`
	ExParStatus        string                   `json:"ExParStatus"`
	ExtraParBlocks     int64                    `json:"ExtraParBlocks"`
	FailedArticles     int64                    `json:"FailedArticles"`
	FileCount          int64                    `json:"FileCount"`
	FileSizeHi         int64                    `json:"FileSizeHi"`
	FileSizeLo         int64                    `json:"FileSizeLo"`
	FileSizeMB         int64                    `json:"FileSizeMB"`
	FinalDir           string                   `json:"FinalDir"`
	FirstID            int64                    `json:"FirstID"`
	Health             int64                    `json:"Health"`
	Kind               string                   `json:"Kind"`
	LastID             int64                    `json:"LastID"`
	Log                []interface{}            `json:"Log"`
	MarkStatus         string                   `json:"MarkStatus"`
	MaxPostTime        int64                    `json:"MaxPostTime"`
	MaxPriority        int64                    `json:"MaxPriority"`
	MessageCount       int64                    `json:"MessageCount"`
	MinPostTime        int64                    `json:"MinPostTime"`
	MinPriority        int64                    `json:"MinPriority"`
	MoveStatus         string                   `json:"MoveStatus"`
	NZBFilename        string                   `json:"NZBFilename"`
	Nzbid              int64                    `json:"NZBID"`
	NZBName            string                   `json:"NZBName"`
	NZBNicename        string                   `json:"NZBNicename"`
	ParStatus          string                   `json:"ParStatus"`
	ParTimeSec         int64                    `json:"ParTimeSec"`
	Parameters         []NzbgetParameters `json:"Parameters"`
	PausedSizeHi       int64                    `json:"PausedSizeHi"`
	PausedSizeLo       int64                    `json:"PausedSizeLo"`
	PausedSizeMB       int64                    `json:"PausedSizeMB"`
	PostInfoText       string                   `json:"PostInfoText"`
	PostStageProgress  int64                    `json:"PostStageProgress"`
	PostStageTimeSec   int64                    `json:"PostStageTimeSec"`
	PostTotalTimeSec   int64                    `json:"PostTotalTimeSec"`
	RemainingFileCount int64                    `json:"RemainingFileCount"`
	RemainingParCount  int64                    `json:"RemainingParCount"`
	RemainingSizeHi    int64                    `json:"RemainingSizeHi"`
	RemainingSizeLo    int64                    `json:"RemainingSizeLo"`
	RemainingSizeMB    int64                    `json:"RemainingSizeMB"`
	RepairTimeSec      int64                    `json:"RepairTimeSec"`
	ScriptStatus       string                   `json:"ScriptStatus"`
	ScriptStatuses     []interface{}            `json:"ScriptStatuses"`
	ServerStats        []NzbgetServerStats `json:"ServerStats"`
	Status             string                   `json:"Status"`
	SuccessArticles    int64                    `json:"SuccessArticles"`
	TotalArticles      int64                    `json:"TotalArticles"`
	URL                string                   `json:"URL"`
	UnpackStatus       string                   `json:"UnpackStatus"`
	UnpackTimeSec      int64                    `json:"UnpackTimeSec"`
	URLStatus          string                   `json:"UrlStatus"`
}

type NzbgetServerStats struct {
	FailedArticles  int64 `json:"FailedArticles"`
	ServerID        int64 `json:"ServerID"`
	SuccessArticles int64 `json:"SuccessArticles"`
}

type NzbgetParameters struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

var (
	articleCacheHi = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_article_cache_hi",
	})

	articleCacheLo = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_article_cache_lo",
	})

	articleCacheMB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_article_cache_mb",
	})

	averageDownloadRate = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_average_download_rate",
	})

	daySizeHi = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_day_size_hi",
	})

	daySizeLo = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_day_size_lo",
	})

	daySizeMB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_day_size_mb",
	})

	download2Paused = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_download_2_paused",
	})

	downloadLimit = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_download_limit",
	})

	downloadPaused = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_download_paused",
	})

	downloadRate = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_download_rate",
	})

	downloadTimeSec = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_download_time_sec",
	})

	downloadedSizeHi = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_downloaded_size_hi",
	})

	downloadedSizeLo = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_downloaded_size_lo",
	})

	downloadedSizeMB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_downloaded_size_mb",
	})

	feedActive = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_feed_active",
	})

	forcedSizeHi = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_forced_size_hi",
	})

	forcedSizeLo = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_forced_size_lo",
	})

	forcedSizeMB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_forced_size_mb",
	})

	freeDiskSpaceHi = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_free_disk_space_hi",
	})

	freeDiskSpaceLo = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_free_disk_space_lo",
	})

	freeDiskSpaceMB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_free_disk_space_mb",
	})

	monthSizeHi = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_month_size_hi",
	})

	monthSizeLo = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_month_size_lo",
	})

	monthSizeMB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_month_size_mb",
	})

	newsServers = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nzb_news_server_active",
	}, []string{"id"})

	parJobCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_par_job_count",
	})

	postJobCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_post_job_count",
	})

	postPaused = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_post_paused",
	})

	queueScriptCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_queue_script_count",
	})

	remainingSizeHi = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_remaining_size_hi",
	})

	remainingSizeLo = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_remaining_size_lo",
	})

	remainingSizeMB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_remaining_size_mb",
	})

	resumeTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_resume_time",
	})

	scanPaused = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_scan_paused",
	})

	serverPaused = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_server_paused",
	})

	serverStandBy = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_server_standby",
	})

	serverTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_server_time",
	})

	threadCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_thread_count",
	})

	upTimeSec = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_uptime_sec",
	})

	urlCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nzb_url_count",
	})

	downloadStatuses = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nzb_download_statuses",
	}, []string{"status"})
)

func boolToFloat(b bool) float64 {
	if b {
		return float64(1)
	}
	return float64(0)
}

func collectQueue() {
	resp, err := http.Get(fmt.Sprintf("%s/jsonrpc/listgroups", os.Getenv("NZBGET_ADDRESS")))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	queue := &NzbgetQueue{}

	if err := json.NewDecoder(resp.Body).Decode(queue); err != nil {
		log.Fatal(err)
	}

	statuses := map[string]float64{}

	for _, queued := range queue.Result {
		statuses[queued.Status] += 1
		if queued.CriticalHealth != 1000 {
			statuses["UNHEALTHY"] += 1
		}
	}

	for status, count := range statuses {
		downloadStatuses.With(prometheus.Labels{
			"status": status,
		}).Set(count)
	}
}

func collectStatus() {
	resp, err := http.Get(fmt.Sprintf("%s/jsonrpc/status", os.Getenv("NZBGET_ADDRESS")))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	status := &NzbgetStatus{}

	if err := json.NewDecoder(resp.Body).Decode(status); err != nil {
		log.Fatal(err)
	}

	articleCacheHi.Set(float64(status.Result.ArticleCacheHi))
	articleCacheLo.Set(float64(status.Result.ArticleCacheLo))
	articleCacheMB.Set(float64(status.Result.ArticleCacheMB))
	averageDownloadRate.Set(float64(status.Result.AverageDownloadRate))
	daySizeHi.Set(float64(status.Result.DaySizeHi))
	daySizeLo.Set(float64(status.Result.DaySizeLo))
	daySizeMB.Set(float64(status.Result.DaySizeMB))
	download2Paused.Set(boolToFloat(status.Result.Download2Paused))
	downloadLimit.Set(float64(status.Result.DownloadLimit))
	downloadPaused.Set(boolToFloat(status.Result.DownloadPaused))
	downloadRate.Set(float64(status.Result.DownloadRate))
	downloadTimeSec.Set(float64(status.Result.DownloadTimeSec))
	downloadedSizeHi.Set(float64(status.Result.DownloadedSizeHi))
	downloadedSizeLo.Set(float64(status.Result.DownloadedSizeLo))
	downloadedSizeMB.Set(float64(status.Result.DownloadedSizeMB))
	feedActive.Set(boolToFloat(status.Result.FeedActive))
	forcedSizeHi.Set(float64(status.Result.ForcedSizeHi))
	forcedSizeLo.Set(float64(status.Result.ForcedSizeLo))
	forcedSizeMB.Set(float64(status.Result.ForcedSizeMB))
	freeDiskSpaceHi.Set(float64(status.Result.FreeDiskSpaceHi))
	freeDiskSpaceLo.Set(float64(status.Result.FreeDiskSpaceLo))
	freeDiskSpaceMB.Set(float64(status.Result.FreeDiskSpaceMB))
	monthSizeHi.Set(float64(status.Result.MonthSizeHi))
	monthSizeLo.Set(float64(status.Result.MonthSizeLo))
	monthSizeMB.Set(float64(status.Result.MonthSizeMB))

	for _, newsServer := range status.Result.NewsServers {
		newsServers.With(prometheus.Labels{
			"id": fmt.Sprintf("%d", newsServer.ID),
		}).Set(boolToFloat(newsServer.Active))
	}

	parJobCount.Set(float64(status.Result.ParJobCount))
	postJobCount.Set(float64(status.Result.PostJobCount))
	postPaused.Set(boolToFloat(status.Result.PostPaused))
	queueScriptCount.Set(float64(status.Result.QueueScriptCount))
	remainingSizeHi.Set(float64(status.Result.RemainingSizeHi))
	remainingSizeLo.Set(float64(status.Result.RemainingSizeLo))
	remainingSizeMB.Set(float64(status.Result.RemainingSizeMB))
	resumeTime.Set(float64(status.Result.ResumeTime))
	scanPaused.Set(boolToFloat(status.Result.ScanPaused))
	serverPaused.Set(boolToFloat(status.Result.ServerPaused))
	serverStandBy.Set(boolToFloat(status.Result.ServerStandBy))
	serverTime.Set(float64(status.Result.ServerTime))
	threadCount.Set(float64(status.Result.ThreadCount))
	upTimeSec.Set(float64(status.Result.UpTimeSec))
	urlCount.Set(float64(status.Result.URLCount))
}

func main() {
	promHandler := promhttp.Handler()

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		collectStatus()
		collectQueue()
		promHandler.ServeHTTP(w, r)
	})

	http.ListenAndServe(":2112", nil)
}
