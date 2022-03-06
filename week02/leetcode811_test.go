package week02

import (
	"strings"
	"testing"
)

//["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
//["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
func TestSubdomainVisits(t *testing.T) {
	source := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	counts := subdomainVisitsV2(source)
	for s, v := range counts {
		println(s, " === ", v)
	}
}

func TestGetAllDomains(t *testing.T) {
	source := []string{"www", "will", "me", "com"}
	result := getAllDomains(source)
	t.Log(result)
}

func TestGetAllDomainsV2(t *testing.T) {
	source := "www.will.me.com"
	//result := getAllDomains(source)
	result := []string{source}
	for {
		index := strings.Index(source, ".")
		if index < 0 {
			break
		}
		source = source[index+1:]
		result = append(result, source)
		println(source)
	}
	t.Log(result)
}

func TestGetDomainAndCount(t *testing.T) {
	source := "900 google.mail.com"
	count, i, err := getDomainAndCount(source)
	if err != nil {
		println(err)
	}
	println(count)
	for _, v := range i {
		println(v)
	}
}
