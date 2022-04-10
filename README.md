# MAS-regulated Financial Institutions

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

[![GitHub license](https://img.shields.io/badge/LICENSE-BSD--3--CLAUSE-GREEN?style=for-the-badge)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliotwutingfeng/MASRegulatedFinancialInstitutions?style=for-the-badge)](https://goreportcard.com/report/github.com/elliotwutingfeng/MASRegulatedFinancialInstitutions)

Generate a machine-readable `.txt` allowlist of websites belonging to financial institutions regulated by the [Monetary Authority of Singapore (MAS)](https://eservices.mas.gov.sg/fid/institution).

**Disclaimer:** _This project is not sponsored, endorsed, or otherwise affiliated with the Monetary Authority of Singapore._

## Requirements

-   Go >= 1.18

## Setup instructions

`git clone` and `cd` into the project directory, then run the following

```bash
go mod init dummy && go mod tidy
```

## Usage

```bash
go run scraper.go
```

## Libraries/Frameworks used

-   [fasttld](https://github.com/elliotwutingfeng/go-fasttld)
-   [soup](https://github.com/anaskhan96/soup)
