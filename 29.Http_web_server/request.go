type Request struct {
	//HTTP 요청 메드 정보를 갖고 있음
	//("GET", "POST", "PUT", "DELETE")
	Method string

	//HTTP 요청을 보낸 URL 정보를 담고 있음
	//URL 정보를 이용해서 URL에 포함된 데이터를 쿼리해올 수 있음
	URL *url.URL

	// HTTP 프로토콜 버전 정보
	// HTTP 1.0인지 2.0인지 확인
	Proto string // "HTTP/1.0"
	ProtoMajor int //1
	ProtoMinor int //0

	//Header 요청 정보를 담고 있음 (맵 형태로)
	Header Header 

	//HTTP 요청의 실제 데이터를 담고 있는 body 정보
	Body io.ReadCloser
}
