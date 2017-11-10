package demoinfocs

import (
	"fmt"
	"sync"

	"github.com/ghostanalysis/demoinfocs-golang/msg"
	"github.com/gogo/protobuf/proto"
)

var packetEntitiesPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return new(msg.CSVCMsg_PacketEntities)
	},
}

var gameEventPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return new(msg.CSVCMsg_GameEvent)
	},
}

var tickPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return new(msg.CNETMsg_Tick)
	},
}

var byteSlicePool sync.Pool = sync.Pool{
	New: func() interface{} {
		s := make([]byte, 0, 256)
		return &s
	},
}

func (p *Parser) parsePacket() {
	for !p.bitReader.ChunkFinished() {
		cmd := int(p.bitReader.ReadVarInt32())
		size := int(p.bitReader.ReadVarInt32())

		p.bitReader.BeginChunk(size << 3)
		var m proto.Message
		switch cmd {
		case int(msg.SVC_Messages_svc_PacketEntities):
			m = packetEntitiesPool.Get().(*msg.CSVCMsg_PacketEntities)
			defer packetEntitiesPool.Put(m)

		case int(msg.SVC_Messages_svc_GameEventList):
			m = new(msg.CSVCMsg_GameEventList)

		case int(msg.SVC_Messages_svc_GameEvent):
			m = gameEventPool.Get().(*msg.CSVCMsg_GameEvent)
			defer gameEventPool.Put(m)

		case int(msg.SVC_Messages_svc_CreateStringTable):
			m = new(msg.CSVCMsg_CreateStringTable)

		case int(msg.SVC_Messages_svc_UpdateStringTable):
			m = new(msg.CSVCMsg_UpdateStringTable)

		case int(msg.NET_Messages_net_Tick):
			m = tickPool.Get().(*msg.CNETMsg_Tick)
			defer tickPool.Put(m)

		case int(msg.SVC_Messages_svc_UserMessage):
			m = new(msg.CSVCMsg_UserMessage)

		default:
			// We don't care about anything else for now
			p.bitReader.EndChunk()
			continue
		}

		b := byteSlicePool.Get().(*[]byte)
		p.bitReader.ReadBytesInto(b, size)

		if proto.Unmarshal(*b, m) != nil {
			// TODO: Don't crash here, happens with demos that work in gotv
			panic(fmt.Sprintf("Failed to unmarshal cmd %d", cmd))
		}
		p.msgQueue <- m

		// Reset to 0 length and pool
		*b = (*b)[:0]
		byteSlicePool.Put(b)

		p.bitReader.EndChunk()
	}

	// Make sure the created events are consumed so they can be pooled
	p.msgDispatcher.SyncQueues(p.msgQueue)
}
