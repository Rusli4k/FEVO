package rest

import (
	"fmt"
	"net/http"

	"github.com/rusli4k/fevo/pkg/parser"
)

// UploadTransaction will handle transactions upload from
// CSV file to repository.
func (th TAHandler) UploadTransactions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	r.ParseMultipartForm(32 << 20) // limit your max input length 32Mb
		file, _, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		ts, err := parser.CSVToTransactions(file)
		if err != nil {
			WriteJSONResponse(w, http.StatusUnprocessableEntity, Response{
				Message: "Error occurred while processing the file",
				Details: err.Error()})
			return
		}

		for _, v := range ts {
			if err := th.usecase.UploadTr(v); err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: "Error while adding to db: ",
					Details: err.Error()})
				return
			}
		}

		WriteJSONResponse(w, http.StatusOK, Response{
			Message: "Request  processed with no errors.",
			Details: fmt.Sprint("Num of added rows:", len(ts))})
	})
}

func (ta TAHandler) GetTransactions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

// 		var user entities.User

// 		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
// 			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)
// 			uh.logger.Errorf("Failed decoding JSON from request %+v: %+v", req, err)

// 			return
// 		}

// 		defer func() {
// 			if err := req.Body.Close(); err != nil {
// 				uh.logger.Warnf("Failed closing request %+v: %+v", req, err)
// 			}
// 		}()

// 		if err := user.Validate(); err != nil {
// 			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)
// 			uh.logger.Errorf("Failed validating user: %+v", err)

// 			return
// 		}

// 		id, err := uh.usecase.SignUp(user)
// 		if err != nil {
// 			uh.logger.Errorf("Failed creating user: %+v", err)

// 			if errors.Is(err, globals.ErrDuplicateEmail) {
// 				WriteJSONResponse(w, http.StatusConflict, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)

// 				return
// 			}

// 			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr}, uh.logger)

// 			return
// 		}

// 		WriteJSONResponse(w, http.StatusCreated, CreateUserResponse{ID: id}, uh.logger)
// 		uh.logger.Debugw("User successfully created", "ID", id)
// 	})
// }

// // GetUserByID will handle user search.
// func (ta TAHandler) GetTransaction() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		id := mux.Vars(req)["id"]

// 		if _, err := uuid.Parse(id); err != nil {
// 			uh.logger.Warnw("Invalid UUID", "ID", id)
// 			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)

// 			return
// 		}

// 		entUser, err := uh.usecase.GetUser(id)
// 		if err != nil {
// 			if errors.Is(err, globals.ErrNotFound) {
// 				uh.logger.Debugw("No user found.", "ID", id)
// 				WriteJSONResponse(w, http.StatusNotFound, Response{Message: MsgNotFound, Details: err.Error()}, uh.logger)

// 				return
// 			}
// 			uh.logger.Errorw("Internal error while searching user.", "ID", id, "error", err.Error())
// 			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr}, uh.logger)

// 			return
// 		}

// 		user := User{
// 			ID:        entUser.ID,
// 			FirstName: entUser.FirstName,
// 			LastName:  entUser.LastName,
// 			Email:     entUser.Email,
// 			CreatedAt: entUser.CreatedAt,
// 		}
// 		WriteJSONResponse(w, http.StatusOK, user, uh.logger)
// 	})
// }

// // // GetUsers retrieves all entities.User by given parameters.
// // func (uh UserHandler) GetUsers() http.Handler {
// // 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// // 		var params = map[string]string{
// // 			sort:   firstName + "," + lastName,
// // 			offset: "",
// // 			limit:  "",
// // 		}

// // 		for k, v := range req.URL.Query() {
// // 			params[k] = strings.Join(v, "")
// // 		}

// // 		users, err := uh.usecase.Fetch(params[offset], params[limit], params[sort])
// // 		if err != nil {
// // 			uh.logger.Errorf("Failed fetching users from repository: %+v", err)
// // 			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: "could not fetch users"}, uh.logger)

// // 			return
// // 		}

// // 		res := struct {
// // 			Results []entities.User `json:"results"`
// // 		}{Results: users}

// // 		WriteJSONResponse(w, http.StatusOK, res, uh.logger)
// // 	})

// // handler to get file from POST request?
// func ReceiveFile(w http.ResponseWriter, r *http.Request) {

//     r.ParseMultipartForm(32 << 20) // limit your max input length!
//     var buf bytes.Buffer
//     // in your case file would be fileupload
//     file, header, err := r.FormFile("file")
//     if err != nil {
//         panic(err)
//     }
//     defer file.Close()
//     name := strings.Split(header.Filename, ".")
//     fmt.Printf("File name %s\n", name[0])
//     // Copy the file data to my buffer
//     io.Copy(&buf, file)
//     // do something with the contents...
//     // I normally have a struct defined and unmarshal into a struct, but this will
//     // work as an example
//     contents := buf.String()
//     fmt.Println(contents)
//     // I reset the buffer in case I want to use it again
//     // reduces memory allocations in more intense projects
//     buf.Reset()
//     // do something else
//     // etc write header
//     return
// }

// // file POST handler?
// func SendPostRequest(url string, filename string) (string, []byte) {
//     api_key := ReadAPIKey("../.api_key")
//     client := &http.Client{}
//     data, err := os.Open(filename)
//     if err != nil {
//         log.Fatal(err)
//     }
//     req, err := http.NewRequest("POST", url, data)
//     if err != nil {
//         log.Fatal(err)
//     }
//     req.SetBasicAuth("api", api_key)
//     resp, err := client.Do(req)
//     if err != nil {
//         log.Fatal(err)
//     }
//     content, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         log.Fatal(err)
//     }
//     return resp.Status, content
// }

// func main() {
//     status, content := SendPostRequest("https://api.example.com/upload", "test.jpg")
//     fmt.Println(status)
//     fmt.Println(string(content))
// }

// // how to handle POST request with file?
// import (
//     "bytes"
//     "io"
//     "mime/multipart"
//     "net/http"
//     "path/filepath"
// )

// // content is a struct which contains a file's name, its type and its data.
// type content struct {
//     fname string
//     ftype string
//     fdata []byte
// }

// func sendPostRequest(url string, files ...content) ([]byte, error) {
//     var (
//         buf = new(bytes.Buffer)
//         w   = multipart.NewWriter(buf)
//     )

//     for _, f := range files {
//         part, err := w.CreateFormFile(f.ftype, filepath.Base(f.fname))
//         if err != nil {
//             return []byte{}, err
//         }
//         part.Write(f.fdata)
//     }

//     w.Close()

//     req, err := http.NewRequest("POST", url, buf)
//     if err != nil {
//         return []byte{}, err
//     }
//     req.Header.Add("Content-Type", w.FormDataContentType())

//     client := &http.Client{}
//     res, err := client.Do(req)
//     if err != nil {
//         return []byte{}, err
//     }
//     defer res.Body.Close()

//     cnt, err := io.ReadAll(res.Body)
//     if err != nil {
//         return []byte{}, err
//     }
//     return cnt, nil
// }
