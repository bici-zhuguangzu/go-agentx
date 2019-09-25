// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import "github.com/bici-zhuguangzu/go-agentx/pdu"

type request struct {
	headerPacket *pdu.HeaderPacket
	responseChan chan *pdu.HeaderPacket
}
