package source

import (
	"encoding/binary"
	"fmt"
	"strings"
	"testing"

	"github.com/Fantom-foundation/go-lachesis/gossip"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

func TestPrintRLP(t *testing.T) {
	t.SkipNow()

	t.Run("gossip.PeerProgress", func(t *testing.T) {
		var progress gossip.PeerProgress
		progress.Epoch = 1
		progress.NumOfBlocks = 3

		bb, err := rlp.EncodeToBytes(&progress)
		if err != nil {
			panic(err)
		}

		printRLP("", bb)
	})

	t.Run("unknown struct", func(t *testing.T) {
		const dump = `82 01 00 d9 01 d0 f8 d7 82 0d 5c 83 24 be 1c f8 ad 2b 82 ec 23 81 fa f8 a5 a0 00 00 0d 5c 00 00
 0a 53 bd dd f5 6d 79 c9 c7 c8 18 e4 7e 4e 87 cd ac 21 58 19 5b 61 e7 73 93 28 a0 0d 21 60 67 06
 75 08 d8 71 52 02 b6 0c 38 2b 4b 29 b8 35 77 a6 b3 da fc 27 3d 22 e3 11 42 68 6d 09 0e 9b a4 3a
 6f f2 76 b2 6d d5 5e b5 96 99 82 dd 8c 48 26 58 4d 60 17 a0 00 09 63 60 6d 1e 33 9f e7 44 00 1e
 cc a4 0f dc c8 66 e1 37 ea 94 46 af d7 91 4d bb ff 11 21 60 6e 33 65 ab 19 db 4e b3 86 3a e0 02
 d0 3c a8 e0 4a 80 88 cb 99 c9 3b 02 ee 11 21 60 50 cf bc 7b 5d 1e 8a 7e c6 60 ba df 44 a5 46 bf
 55 41 e7 0d 74 d4 2b 4b 38 00 00 00 00 00 00 00 6a 1c 91 30 5b a1 1b ca dd b5 6a 96 8b c6 f2 d1`

		hex := strings.ReplaceAll(strings.ReplaceAll(dump, " ", ""), "\n", "")
		raw := common.Hex2Bytes(hex)

		msg := binary.LittleEndian.Uint32(raw[:4])
		size := int(binary.LittleEndian.Uint16(raw[4:6]))
		fmt.Printf("Msg: %d\nSize: %d\nBody:\n", msg, size)

		body := raw[6:]
		size = len(body)
		body = append(body, make([]byte, size-len(body))...)

		printRLP("", body)
	})
}

func printRLP(ident string, bb []byte) (rest []byte) {
	var (
		kind rlp.Kind
		val  []byte
		err  error
	)
	for len(bb) > 0 {
		kind, val, bb, err = rlp.Split(bb)
		if err != nil {
			panic(err)
			return bb
		}
		switch kind {
		case rlp.Byte:
			fmt.Printf("%s raw num: %v\n", ident, val)
		case rlp.String:
			switch len(val) {
			case common.HashLength:
				fmt.Printf("%s hash: %s\n", ident, common.BytesToHash(val).Hex())
			case common.AddressLength:
				fmt.Printf("%s addr: %s\n", ident, common.BytesToAddress(val).Hex())
			default:
				fmt.Printf("%s raw data: %v\n", ident, val)
			}
		case rlp.List:
			printRLP(ident+"_", val)
		default:
			panic(fmt.Sprintf("Unknown(%d)", kind))
		}
	}
	return bb
}
