package config
//
//import (
//	"github.com/hashicorp/go-plugin"
//	"net/rpc"
//)
//
//
//func (s *NotifierRPCServer) Notify(mapArgs map[string]interface{}, resp *string) error {
//	args := mapArgs["data"].(TaskEvent)
//	*resp = s.Impl.Notify(args)
//	return nil
//}
//
//func (g *NotifierRPCClient) Notify(event TaskEvent) string {
//	var resp string
//	err := g.client.Call("Plugin.Notify", map[string]interface{}{
//		"data":   event,
//	}, &resp)
//	if err != nil {
//		// You usually want your interfaces to return errors. If they don't,
//		// there isn't much other choice here.
//		panic(err)
//	}
//
//	return resp
//}
//
//type NotifierRPCClient struct {
//	client *rpc.Client
//}
//
//type NotifierRPCServer struct {
//	Impl Notifier
//}
//
//type NotifierPlugin struct {
//	Impl Notifier
//}
//
//func (r *NotifierPlugin) Server(broker *plugin.MuxBroker) (interface{}, error){
//	return &NotifierRPCServer{Impl: r.Impl}, nil
//}
//
//func (r *NotifierPlugin) Client (broker *plugin.MuxBroker, c *rpc.Client) (interface{}, error){
//	return &NotifierRPCClient{client: c}, nil
//}