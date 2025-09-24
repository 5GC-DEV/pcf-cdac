// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package notifyevent

import (
	"github.com/5GC-DEV/event"
	"github.com/omec-project/openapi/models"
	"github.com/omec-project/pcf/logger"
)

var notifyDispatcher *event.Dispatcher

func RegisterNotifyDispatcher() error {
	notifyDispatcher = event.NewDispatcher()
	if err := notifyDispatcher.Register(NotifyListener{},
		SendSMpolicyUpdateNotifyEventName,
		SendSMpolicyTerminationNotifyEventName); err != nil {
		return err
	}
	logger.NotifyEventLog.Debugf("Event handlers registered: %s, %s",
		SendSMpolicyUpdateNotifyEventName,
		SendSMpolicyTerminationNotifyEventName)
	return nil
}

func DispatchSendSMPolicyUpdateNotifyEvent(uri string, request *models.SmPolicyNotification) {
	logger.NotifyEventLog.Infof("DispatchSendSMPolicyUpdateNotifyEvent triggered")
	logger.NotifyEventLog.Debugf("Target URI: %s", uri)

	if notifyDispatcher == nil {
		logger.NotifyEventLog.Errorf("notifyDispatcher is nil")
		return
	}

	logger.NotifyEventLog.Debugf("Sending SM Policy Update Notify Event to dispatcher")

	err := notifyDispatcher.Dispatch(SendSMpolicyUpdateNotifyEventName, SendSMpolicyUpdateNotifyEvent{
		uri:     uri,
		request: request,
	})

	if err != nil {
		logger.NotifyEventLog.Errorf("Failed to dispatch SM Policy Update Notify Event: %v", err)
	} else {
		logger.NotifyEventLog.Infof("Successfully dispatched SM Policy Update Notify Event")
	}
}

func DispatchSendSMPolicyTerminationNotifyEvent(uri string, request *models.TerminationNotification) {
	if notifyDispatcher == nil {
		logger.NotifyEventLog.Errorf("notifyDispatcher is nil")
	}
	err := notifyDispatcher.Dispatch(SendSMpolicyTerminationNotifyEventName, SendSMpolicyTerminationNotifyEvent{
		uri:     uri,
		request: request,
	})
	if err != nil {
		logger.NotifyEventLog.Errorln(err)
	}
}
