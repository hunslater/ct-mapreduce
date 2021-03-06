/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package config

import (
	"flag"
	"github.com/vharitonsky/iniflags"
)

type CTConfig struct {
	LogUrlList        *string
	Verbose           *bool
	CertPath          *string
	Offset            *uint64
	Limit             *uint64
	NumThreads        *int
	RunForever        *bool
	PollingDelay      *int
	IssuerCNFilter    *string
	LogExpiredEntries *bool
	AWSS3Bucket       *string
}

func NewCTConfig() *CTConfig {
	ret := &CTConfig{
		LogUrlList:        flag.String("logList", "", "URLs of the CT Logs, comma delimited"),
		Verbose:           flag.Bool("v", false, "Give verbose output"),
		CertPath:          flag.String("certPath", "", "Path under which to store full DER-encoded certificates"),
		Offset:            flag.Uint64("offset", 0, "offset from the beginning"),
		Limit:             flag.Uint64("limit", 0, "limit processing to this many entries"),
		NumThreads:        flag.Int("numThreads", 1, "Use this many threads per CPU"),
		RunForever:        flag.Bool("forever", false, "Run forever"),
		PollingDelay:      flag.Int("pollingDelay", 10, "Wait this many minutes between polls"),
		IssuerCNFilter:    flag.String("issuerCNList", "", "Prefixes to match for CNs for permitted issuers, comma delimited"),
		LogExpiredEntries: flag.Bool("logExpiredEntries", false, "Add expired entries to the database"),
		AWSS3Bucket:         flag.String("awsS3Bucket", "", "AWS S3 Path"),
	}

	iniflags.Parse()
	return ret
}

func (c *CTConfig) Usage() {
	flag.Usage()
}
