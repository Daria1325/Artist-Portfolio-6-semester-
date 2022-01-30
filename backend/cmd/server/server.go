package server

import (
	"github.com/daria/Portfolio/backend/cmd/database"
)

type Server struct {
	Repo *database.Repo
}

//func (d *Server) GetPictures(ctx context.Context) (string, error) {
//	ports, err := d.Repo.GetPictures()
//	if err != nil {
//		return nil, err
//	}
//	w, err := json.Marshal(ports)
//	if err != nil {
//		return nil, err
//	}
//	return &api.GetPortsResponse{List: string(w)}, nil
//}
//func (d *Server) GetPort(ctx context.Context, req *api.GetPortRequest) (*api.GetPortResponse, error) {
//	ports, err := d.Repo.GetPorts()
//	if err != nil {
//		return nil, err
//	}
//
//	for _, item := range ports {
//		if strconv.Itoa(item.ID) == req.Id {
//			w, err := json.Marshal(item)
//			if err != nil {
//				return nil, err
//			}
//			w = []byte(fmt.Sprintf("[%s]", w))
//			return &api.GetPortResponse{Item: string(w)}, nil
//		}
//	}
//	return &api.GetPortResponse{Item: "Not found, check the Id"}, errors.New("Not found")
//}
//
//func (d *Server) UpsertPorts(stream api.Port_UpsertPortsServer) error {
//	var buf []byte
//	updatedId := "Id: "
//	ports, err := d.Repo.GetPorts()
//	if err != nil {
//		return err
//	}
//
//	var portArray []database.Port
//
//	for {
//		recv, err := stream.Recv()
//		if err != nil {
//			if err == io.EOF {
//				goto END
//			}
//
//			err = errors.Wrapf(err,
//				"failed unexpectadely while reading chunks from stream")
//			return err
//		}
//		buf = append(buf,recv.GetContent()...)
//		}
//
//END:
//	err = json.Unmarshal(buf, &portArray)
//	if err != nil {
//		log.Print(err)
//		return err
//	}
//	for _, port := range portArray {
//		isNotInDatabase := true
//		if len(ports) != 0 {
//			for _, item := range ports {
//				if item.ID == port.ID {
//					updatedId = fmt.Sprintf("%s %d", updatedId, item.ID)
//					isNotInDatabase = false
//					err := d.Repo.UpdatePort(port)
//					if err != nil {
//						return err
//					}
//					continue
//				}
//			}
//		}
//
//		if isNotInDatabase {
//			err := d.Repo.AddPort(port)
//			if err != nil {
//				return err
//			}
//		}
//	}
//
//
//	err = stream.SendAndClose(&api.UpsertPortsResponse{
//		Message: updatedId,
//		Code:    api.UploadStatusCode_Ok,
//	})
//	if err != nil {
//		err = errors.Wrapf(err,
//			"failed to send status code")
//		return err
//	}
//	return err
//}
