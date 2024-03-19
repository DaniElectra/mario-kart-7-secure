package nex

import (
	"os"
	"strconv"

	"github.com/PretendoNetwork/mario-kart-7/globals"
	nex_ticket_granting "github.com/PretendoNetwork/mario-kart-7/nex/ticket-granting"
	"github.com/PretendoNetwork/nex-go/constants"
	"github.com/PretendoNetwork/nex-go/types"
	common_ticket_granting "github.com/PretendoNetwork/nex-protocols-common-go/ticket-granting"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-go/ticket-granting"
)

func registerCommonAuthenticationServerProtocols() {
	ticketGrantingProtocol := ticket_granting.NewProtocol()
	globals.AuthenticationEndpoint.RegisterServiceProtocol(ticketGrantingProtocol)
	globals.CommonTicketGrantingProtocol = common_ticket_granting.NewCommonProtocol(ticketGrantingProtocol)
	ticketGrantingProtocol.SetHandlerLoginEx(nex_ticket_granting.LoginEx)
	ticketGrantingProtocol.SetHandlerRequestTicket(nex_ticket_granting.RequestTicket)

	port, _ := strconv.Atoi(os.Getenv("PN_MK7_SECURE_SERVER_PORT"))

	secureStationURL := types.NewStationURL("")
	secureStationURL.SetURLType(constants.StationURLPRUDPS)
	secureStationURL.SetAddress(os.Getenv("PN_MK7_SECURE_SERVER_HOST"))
	secureStationURL.SetPortNumber(uint16(port))
	secureStationURL.SetConnectionID(1)
	secureStationURL.SetPrincipalID(types.NewPID(2))
	secureStationURL.SetStreamID(1)
	secureStationURL.SetStreamType(constants.StreamTypeRVSecure)
	secureStationURL.SetType(uint8(constants.StationURLFlagPublic))

	globals.CommonTicketGrantingProtocol.SecureStationURL = secureStationURL
	globals.CommonTicketGrantingProtocol.BuildName = types.NewString(serverBuildString)
	globals.CommonTicketGrantingProtocol.SecureServerAccount = globals.SecureServerAccount
}
