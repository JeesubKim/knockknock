package main

type Col struct {
	content string
}

// Word 에 Mapping 되는 Col이 있고 이게 최소단위
// Word마다 Col 갯수가 다르므로 이게 연결
type Word struct {
	col_id []int64
}

//Word의 CRUD로 하고 이게 Col로 바뀌고 DB로 넣는다
func InitWord(router http.Handler) {

	router.HandleFunc("GET /api/word/", handleGetAllWords)
	router.HandleFunc("POST /api/word/", handleCreateNewWord)
	router.HandleFunc("GET /api/word/{id}", handleGetWordById)
	router.HandleFunc("UPDATE /api/word/{id}", handleUpdateWord)
	router.HandleFunc("DELETE /api/word/{id}", handleDeleteWord)
}

func handleGetAllWords(w http.ResponseWriter, r *http.Request) {

}

func handleGetWordById(w http.ResponseWriter, r *http.Request) {

}

func handleCreateNewWord(w http.ResponseWriter, r *http.Request) {

}

func handleUpdateWord(w http.ResponseWriter, r *http.Request) {

}

func handleDeleteWord(w http.ResponseWriter, r *http.Request) {

}
