package week02

import (
	"strconv"
	"strings"
)

// ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
// ["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
// 第一步 ：按照空格拆分出域名及其访问次数
// 第二步 ：将域名按分隔符截取后直接和次数一起存入 map 中
// 第三步 ：将 map 转换为输出结果
func subdomainVisitsV2(domains []string) []string {
	m := make(map[string]int)
	for _, v := range domains {
		split := strings.Split(v, " ")
		count, _ := strconv.Atoi(split[0])
		allDomains := getAllDomainsV2(split[1])
		for _, v := range allDomains {
			m[v] += count
		}
	}
	return map2result(m)
}

// in : 900, "google.mail.com"
func getMapForDomainAndCount(count int, s string) {
	var result []string
	for strings.Contains(s, ".") {
		result = append(result, s)
	}

	//m := make(map[string]int)

}

// ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
// ["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
// 第一步 ：按照空格拆分出域名及其访问次数
// 第二步 ：将域名按分隔符拆分后遍历出所有域名组合
// 第三步 ：将所有域名组合作为 key 访问次数作为 value 放入 map 中
// 第四步 ：将 map 转换为输出结果
func subdomainVisitsV1(domains []string) []string {
	counts := make(map[string]int)
	for _, v := range domains {
		count, array, _ := getDomainAndCount(v)
		allDomains := getAllDomains(array)
		for _, v := range allDomains {
			counts[v] += count
		}
	}
	var result []string
	for k, v := range counts {
		result = append(result, strconv.Itoa(v)+" "+k)
	}
	return result
}

func map2result(m map[string]int) []string {
	var result []string
	for k, v := range m {
		result = append(result, strconv.Itoa(v)+" "+k)
	}
	return result
}

// in : "900 google.mail.com"
// out : 900, [google, mail, com], err
func getDomainAndCount(s string) (int, []string, error) {
	split := strings.Split(s, " ")
	count, err := strconv.Atoi(split[0])
	array := strings.Split(split[1], ".")
	return count, array, err
}

// in : "google.mail.com"
// out : "com" "mail.com" "google.mail.com"
func getAllDomains(source []string) []string {
	var sb strings.Builder
	var result []string
	n := len(source)
	for i := n - 1; i >= 0; i-- {
		v := source[i]
		sb.WriteString(v)
		if i != n-1 {
			sb.WriteString(".")
			sb.WriteString(result[len(result)-1])
		}
		result = append(result, sb.String())
		sb.Reset()
	}
	return result
}

// in : "google.mail.com"
// out : "com" "mail.com" "google.mail.com"
func getAllDomainsV2(source string) []string {
	result := []string{source}
	for {
		index := strings.Index(source, ".")
		if index < 0 {
			break
		}
		source = source[index+1:]
		result = append(result, source)
	}
	return result
}
