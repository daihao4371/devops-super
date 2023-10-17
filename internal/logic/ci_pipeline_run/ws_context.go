package ci_pipeline_run

import (
	"bufio"
	"context"
	"devops-super/utility/thirdclients/kubernetes"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type wsContext struct {
	request      *ghttp.Request
	ws           *ghttp.WebSocket
	ctx          context.Context
	cancelFunc   context.CancelFunc
	kubeClient   *kubernetes.Client
	watcher      watch.Interface
	namespace    string
	podName      string
	lastPingTime *gtime.Time
}

// 获取 Pod 日志
func (wsCtx *wsContext) tailLog(logIndex int) error {
	line := int64(100000)
	req := wsCtx.kubeClient.CoreV1().Pods(wsCtx.namespace).GetLogs(wsCtx.podName, &corev1.PodLogOptions{
		Container: fmt.Sprintf("env-%d", logIndex),
		Follow:    true,
		TailLines: &line,
	})
	stream, err := req.Stream(wsCtx.ctx)
	if err != nil {
		return err
	}
	defer stream.Close()
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		if err := wsCtx.ws.WriteMessage(ghttp.WsMsgText, scanner.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

func (wsCtx *wsContext) checkClientClose() {
	var handleClose = func() {
		wsCtx.cancelFunc()
		if wsCtx.watcher != nil {
			wsCtx.watcher.Stop()
		}
		wsCtx.ws.Close()
	}

	for {
		_, _, err := wsCtx.ws.ReadMessage() // 仅用来监听客户端连接关闭
		if err != nil {
			//if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			//	// 连接正常关闭或正在关闭
			//	fmt.Println("连接关闭:", err)
			//} else {
			//	// 连接异常关闭
			//	fmt.Println("连接异常关闭:", err)
			//}
			handleClose()
			break
		}
	}
}
