var ue *pcf_context.UeContext
ueRegisteredFlag := 0 // Initialize the flag

if val, exist := pcfSelf.UePool.Load(request.Supi); exist {
    ue = val.(*pcf_context.UeContext)
    if ue != nil && ue.Registered {
        ueRegisteredFlag = 1
    }
}

if ue == nil {
    problemDetail := util.GetProblemDetail("Supi is not supported in PCF", util.USER_UNKNOWN)
    logger.SMpolicylog.Warnf("Supi[%s] is not supported in PCF", request.Supi)
    return nil, nil, &problemDetail
}

if ueRegisteredFlag == 0 {
    problemDetail := util.GetProblemDetail("UE is not registered in PCF", util.USER_UNKNOWN)
    logger.SMpolicylog.Warnf("UE[%s] is not registered in PCF", ue.Supi)
    return nil, nil, &problemDetail
}

// ✅ Only execute this block if UE is registered
logger.SMpolicylog.Infof(" AM Policy Parameters 1: [%s]", request.AccessType)
logger.SMpolicylog.Infof(" AM Policy Parameters 2: [%s]", request.ServingNetwork)

amPolicy := ue.FindAMPolicy(request.AccessType, request.ServingNetwork)
logger.SMpolicylog.Infof(" AM PolicyFind: [%s]", amPolicy)

if amPolicy == nil {
    problemDetail := util.GetProblemDetail("Can't find corresponding AM Policy", util.POLICY_CONTEXT_DENIED)
    logger.SMpolicylog.Warnln("can not find corresponding AM Policy")
    return nil, nil, &problemDetail
}
==============================================================================================