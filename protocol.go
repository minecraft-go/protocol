package protocol

type (
	State     int
	Direction int
	Version   int
)

const (
	Handshaking State = iota
	Status
	Login
	Play
)

func (state State) String() string {
	switch state {
	case Handshaking:
		return "Handshaking"
	case Status:
		return "Status"
	case Login:
		return "Login"
	case Play:
		return "Play"
	default:
		return "Unknown"
	}
}

const (
	ClientBound Direction = iota
	ServerBound
)

func (direction Direction) String() string {
	switch direction {
	case ClientBound:
		return "ClientBound"
	case ServerBound:
		return "ServerBound"
	default:
		return "Unknown"
	}
}

//goland:noinspection GoSnakeCaseUsage,GoUnusedConst
const (
	Unknown Version = -1
	v1_7_2  Version = 4
	v1_7_6  Version = 5
	V1_8    Version = 47
	V1_9    Version = 107
	V1_9_1  Version = 108
	V1_9_2  Version = 109
	V1_9_3  Version = 110
	V1_10   Version = 210
	V1_11   Version = 315
	V1_11_1 Version = 316
	V1_12   Version = 335
	V1_12_1 Version = 338
	V1_12_2 Version = 340
	V1_13   Version = 393
	V1_13_1 Version = 401
	V1_13_2 Version = 404
	V1_14   Version = 477
	V1_14_1 Version = 480
	V1_14_2 Version = 485
	V1_14_3 Version = 490
	V1_14_4 Version = 498
	V1_15   Version = 573
	V1_15_1 Version = 575
	V1_15_2 Version = 578
	V1_16   Version = 735
	V1_16_1 Version = 736
	V1_16_2 Version = 751
	V1_16_3 Version = 753
	V1_16_4 Version = 754
	V1_17   Version = 755
	V1_17_1 Version = 756
	V1_18   Version = 757
)
