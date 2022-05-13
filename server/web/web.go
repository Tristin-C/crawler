package web

import (
	"context"
	"crawler/conf"
	"crawler/model"
	"crawler/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var ctx = context.Background()

type Handler struct {
	svc    *service.Service
	config *conf.Config
}

func NewWebpServer(c *conf.Config, s *service.Service) *Handler {
	h := &Handler{
		svc:    s,
		config: c,
	}
	return h
}

func Start(c *conf.Config, s *service.Service) (err error) {
	h := NewWebpServer(c, s)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/cron/info", h.CronInfo)
	serveMux.HandleFunc("/cron/batch", h.CronBatch)
	serveMux.HandleFunc("/cron/list", h.CronList)
	serveMux.HandleFunc("/cron_log/list", h.CronLogList)

	//2.设置监听的TCP地址并启动服务
	//参数1：TCP地址(IP+Port)
	//参数2：当设置为nil时表示使用DefaultServeMux，如果指定了则表示使用自定义ServeMux
	fmt.Println("server start with addr", c.Http.Addr)
	err = http.ListenAndServe(c.Http.Addr, serveMux)
	return
}

func (h *Handler) CronBatch(w http.ResponseWriter, r *http.Request) {
	resp := &model.Resp{}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)

	req := &CronAddRequest{}
	_ = json.Unmarshal(reqBody, req)

	res, err := h.svc.BatchCron(ctx, req.ConvertToCronAddReq())
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}
	resp.Data = res
	bs, _ := json.Marshal(resp)
	_, _ = w.Write(bs)
}

func (h *Handler) CronInfo(w http.ResponseWriter, r *http.Request) {
	resp := &model.Resp{}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	idStr := query.Get("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}

	res, err := h.svc.CronInfo(context.Background(), id)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}
	resp.Data = res
	bs, _ := json.Marshal(resp)
	_, _ = w.Write(bs)
}

func (h *Handler) CronList(w http.ResponseWriter, r *http.Request) {
	resp := &model.Resp{}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	pageStr := query.Get("page")
	pageSizeStr := query.Get("pageSize")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}

	res, err := h.svc.CronsList(ctx, page, pageSize)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}
	resp.Data = res
	bs, _ := json.Marshal(resp)
	_, _ = w.Write(bs)
}

func (h *Handler) CronLogList(w http.ResponseWriter, r *http.Request) {
	resp := &model.Resp{}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	pageStr := query.Get("page")
	pageSizeStr := query.Get("pageSize")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}

	res, err := h.svc.CronsLogList(ctx, page, pageSize)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}
	resp.Data = res
	bs, _ := json.Marshal(resp)
	_, _ = w.Write(bs)
}
