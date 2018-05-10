/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(volumeManagerVolumesMetric)
	prometheus.MustRegister(reconstructVolumeErrorsMetric)
}

var (
	volumeManagerVolumesMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "volume_manager_total_volumes",
			Help: "Number of volumes in Volume Manager",
		}, []string{"state"})

	reconstructVolumeErrorsMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "volume_manager_reconstruct_volume_errors_total",
			Help: "Amount of errors on ReconstructVolumeOperation",
		},
		[]string{"volume_name"})
)

// RecordVolumeManagerVolumesMetric registers the amount of volumes, and their respective states, in the Volume Manager.
func RecordVolumeManagerVolumesMetric(state string, number int) {
	volumeManagerVolumesMetric.WithLabelValues(state).Set(float64(number))
}

// RecordRecVolErrorsMetric registers the amount of errors that happened on ReconstructVolumeOperation.
func RecordRecVolErrorsMetric(volName string) {
	reconstructVolumeErrorsMetric.WithLabelValues(volName).Inc()
}
