package main

import (
	"fmt"
	ict "github.com/tdh-foundation/icinga2-go-checktools"
	"os"
	"testing"
)

const (
	Cisco2960Response = `

Port      Name               status       Vlan       Duplex  Speed Type 
Gi1/0/1   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/2   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/3   ** User - Phone ** connected    3          a-full a-1000 10/100/1000BaseTX
Gi1/0/4   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/5   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/6   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/7   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/8   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/9   ** User - Phone ** connected    3          a-full a-1000 10/100/1000BaseTX
Gi1/0/10  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/11  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/12  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/13  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/14  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/15  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/16  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/17  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/18  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/19  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/20  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/21  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/22  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/23  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/24  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/25  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/26  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/27  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/28  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/29  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/30  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/31  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/32  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/33  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/34  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/35  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/36  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/37  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/38  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/39  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/40  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/41  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/42  ** User - Phone ** connected    3          a-half   a-10 10/100/1000BaseTX
Gi1/0/43  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi1/0/44  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/45  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/46  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/47  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi1/0/48  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Te1/0/1                      connected    trunk        full    10G SFP-10GBase-SR
Te1/0/2                      notconnect   1            full    10G Not Present
Gi2/0/1   ** User - Phone ** connected    3          a-full a-1000 10/100/1000BaseTX
Gi2/0/2   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/3   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/4   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/5   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/6   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/7   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/8   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/9   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/10  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/11  ** User - Phone ** connected    3          a-full a-1000 10/100/1000BaseTX
Gi2/0/12  ** User - Phone ** connected    3          a-full a-1000 10/100/1000BaseTX
Gi2/0/13  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/14  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/15  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/16  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/17  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/18  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/19  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/20  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/21  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/22  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/23  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/24  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/25  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/26  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/27  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/28  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/29  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/30  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/31  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/32  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/33  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/34  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/35  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/36  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/37  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/38  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/39  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/40  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/41  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/42  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/43  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi2/0/44  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/45  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/46  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/47  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi2/0/48  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Te2/0/1                      notconnect   1            full    10G Not Present
Te2/0/2                      notconnect   1            full    10G Not Present
Gi3/0/1   ** User - Phone ** connected    3          a-full a-1000 10/100/1000BaseTX
Gi3/0/2   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/3   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/4   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/5   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/6   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/7   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/8   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/9   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/10  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/11  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/12  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/13  ** User - Phone ** connected    3          a-full a-1000 10/100/1000BaseTX
Gi3/0/14  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/15  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/16  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/17  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/18  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/19  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/20  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/21  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/22  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/23  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/24  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/25  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/26  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/27  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/28  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/29  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/30  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/31  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/32  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/33  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/34  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/35  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/36  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/37  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/38  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/39  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/40  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/41  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/42  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/43  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/44  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/45  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/46  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi3/0/47  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi3/0/48  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Te3/0/1                      notconnect   1            full    10G Not Present
Te3/0/2                      notconnect   1            full    10G Not Present
Gi4/0/1   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/2   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/3   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/4   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/5   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/6   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/7   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/8   ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/9   ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/10  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/11  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/12  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/13  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/14  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/15  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/16  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/17  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/18  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/19  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/20  ** User - Phone ** err-disabled 3            auto   auto 10/100/1000BaseTX
Gi4/0/21  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/22  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/23  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/24  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/25  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/26  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/27  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/28  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/29  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/30  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/31  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/32  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/33  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/34  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/35  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/36  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/37  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/38  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/39  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/40  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/41  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/42  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/43  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/44  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/45  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Gi4/0/46  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/47  ** User - Phone ** connected    3          a-full  a-100 10/100/1000BaseTX
Gi4/0/48  ** User - Phone ** notconnect   3            auto   auto 10/100/1000BaseTX
Te4/0/1                      notconnect   1            full    10G Not Present
Te4/0/2                      connected    trunk        full    10G SFP-10GBase-SR
Po1                          connected    trunk      a-full    10G 
Po2       ** User - Phone ** notconnect   unassigned   auto   auto 
Fa0                          disabled     routed       auto   auto 10/100BaseTX`

	CiscoNXResponse = `
--------------------------------------------------------------------------------
Port          Name               status    Vlan      Duplex  Speed   Type
--------------------------------------------------------------------------------
mgmt0         --                 connected routed    full    1000    --         
Eth1/1        * Vlan 7 PBX11 *   notconnec 7         full    1000    1000base-T 
Eth1/2        CHLSSW31           connected trunk     full    10G     10Gbase-SR-S
Eth1/3        CHLSSW32           connected trunk     full    10G     10Gbase-SR-S
Eth1/4        CHLSSW41           connected trunk     full    10G     10Gbase-SR-S
Eth1/5        CHLSSW42           connected trunk     full    10G     10Gbase-SR-S
Eth1/6        TDHESX76           connected trunk     full    10G     SFP-H10GB-CU2M
Eth1/7        TDHESX77           connected trunk     full    10G     SFP-H10GB-CU2M
Eth1/8        TDHESX78           connected trunk     full    10G     SFP-H10GB-CU2M
Eth1/9        TDHBKP02           connected 2         full    1000    1000base-T 
Eth1/10       TDHBKP02           connected 2         full    1000    1000base-T 
Eth1/11       * ASA Transit *    connected 4         full    1000    1000base-T 
Eth1/12       * ASA Wifi Guest e connected trunk     full    1000    1000base-T 
Eth1/13       Synology TDHNAS82  notconnec 2         full    1000    1000base-T 
Eth1/14       Aruba Controller   connected trunk     full    1000    1000base-T 
Eth1/15       Aruba Controller   connected trunk     full    1000    1000base-T 
Eth1/16       * TDHNAS05 mgmt *  connected 2         full    1000    1000base-T 
Eth1/17       * TDHDC03 *        notconnec 2         full    1000    1000base-T 
Eth1/18       TDHESX76 NFS       connected trunk     full    10G     SFP-H10GB-CU5M
Eth1/19       TDHESX77 NFS       connected trunk     full    10G     SFP-H10GB-CU5M
Eth1/20       TDHESX78 NFS       connected trunk     full    10G     SFP-H10GB-CU5M
Eth1/21       TDHESX76 Idrac     connected 2         full    1000    1000base-T 
Eth1/22       TDHESX77 Idrac     connected 2         full    1000    1000base-T 
Eth1/23       TDHESX78 Idrac     connected 2         full    1000    1000base-T 
Eth1/24       TDHNAS5 Mgmt       connected 100       full    1000    1000base-T 
Eth1/25       ++ Port H/S - CDB/ xcvrAbsen 1         auto    10G     --         
Eth1/26       TDHESX77 Vswitch0  connected trunk     full    1000    1000base-T 
Eth1/27       TDHESX78 Vswitch0  connected trunk     full    1000    1000base-T 
Eth1/28       --                 xcvrAbsen 1         full    10G     --         
Eth1/29       TDHESX76 Vswitch0  connected trunk     full    1000    1000base-T 
Eth1/30       TDHBKP03           connected 2         full    10G     SFP-H10GB-CU2M
Eth1/31       --                 xcvrAbsen 1         full    10G     --         
Eth1/32       --                 xcvrAbsen 1         full    10G     --         
Eth1/33       --                 xcvrAbsen 1         full    10G     --         
Eth1/34       --                 xcvrAbsen 1         full    10G     --         
Eth1/35       CHLSSW35           notconnec 1         full    1000    1000base-T 
Eth1/36       --                 xcvrAbsen 1         full    10G     --         
Eth1/37       --                 xcvrAbsen 1         full    10G     --         
Eth1/38       --                 xcvrAbsen 1         full    10G     --         
Eth1/39       --                 xcvrAbsen 1         full    10G     --         
Eth1/40       --                 xcvrAbsen 1         full    10G     --         
Eth1/41       TDHESX79 vSwitch0  connected trunk     full    1000    1000base-T 
Eth1/42       TDHESX79 vSwitch1  connected trunk     full    1000    1000base-T 
Eth1/43       --                 xcvrAbsen 1         full    10G     --         
Eth1/44       --                 xcvrAbsen 1         full    10G     --         
Eth1/45       --                 xcvrAbsen 1         full    10G     --         
Eth1/46       * Netapp TDHNAS06  connected trunk     full    10G     SFP-H10GB-CU5M
Eth1/47       * Netapp TDHNAS05  connected trunk     full    10G     SFP-H10GB-CU5M
Eth1/48       --                 xcvrAbsen 1         full    10G     --         
Eth1/49       --                 xcvrAbsen 1         full    40G     --         
Eth1/50       --                 xcvrAbsen 1         full    40G     --         
Eth1/51       --                 xcvrAbsen 1         full    40G     --         
Eth1/52       --                 xcvrAbsen 1         full    40G     --         
Eth1/53       --                 connected trunk     full    40G     QSFP-40G-CR4
Eth1/54       --                 connected trunk     full    40G     QSFP-40G-CR4
Po1           --                 connected trunk     full    40G     --         
Po2           CHLSSW31           connected trunk     full    10G     --         
Po3           CHLSSW32           connected trunk     full    10G     --         
Po4           CHLSSW41           connected trunk     full    10G     --         
Po5           CHLSSW42           connected trunk     full    10G     --         
Po6           TDHESX76           connected trunk     full    10G     --         
Po7           TDHESX77           connected trunk     full    10G     --         
Po8           TDHESX78           connected trunk     full    10G     --         
Po9           TDHBKP02           connected 2         full    1000    --         
Po14          Aruba Controller   connected trunk     full    1000    --         
Po15          TDHESX79 vSwitch0  noOperMem trunk     full    auto    --         
Po16          TDHESX79 vSwitch1  noOperMem trunk     full    auto    --         
Po18          TDHESX76 NFS       connected trunk     full    10G     --         
Po19          TDHESX77 NFS       connected trunk     full    10G     --         
Po20          TDHESX78 NFS       connected trunk     full    10G     --         
Po30          TDHBKP03           connected 2         full    10G     --         
Po46          * Netapp TDHNAS6 N connected trunk     full    10G     --         
Po47          * Netapp TDHNAS5 N connected trunk     full    10G     --         
Vlan1         --                 down      routed    auto    auto    --
Vlan2         --                 connected routed    auto    auto    --
Vlan3         --                 connected routed    auto    auto    --
Vlan4         --                 connected routed    auto    auto    --
Vlan5         --                 connected routed    auto    auto    --
Vlan7         --                 connected routed    auto    auto    --
Vlan100       --                 connected routed    auto    auto    --
`

	Cisco2960IntefacesCount = 203
	CiscoNXInterfacesCount  = 80
)

var sw ict.SwitchInterface

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

// Testing basic parsing if error occurs with mock response we stop test with failed status
func TestCiscoSwitch_ParseInterfaceStatus(t *testing.T) {
	sw = NewCiscoSwitch(params.host)
	if err := sw.ParseInterfaceStatus(Cisco2960Response); err != nil {
		t.Fatalf("Error parsing C2960 response: %s", err)
	}
	if len(sw.Status()) != Cisco2960IntefacesCount {
		t.Errorf("Error want %d interfaces got %d", Cisco2960IntefacesCount, len(sw.Status()))
	}
	fmt.Println(sw.ReturnIcingaResult())

	wantStatus := ict.SwitchInterfaceStatus{Port: "Gi2/0/11",
		Name:   "** User - Phone **",
		Status: "connected",
		Vlan:   "3",
		Duplex: "a-full",
		Speed:  "a-1000",
		Type:   "10/100/1000BaseTX",
	}

	gi2011 := sw.Status()[60]
	if gi2011.Port != wantStatus.Port ||
		gi2011.Name != wantStatus.Name ||
		gi2011.Status != wantStatus.Status ||
		gi2011.Vlan != wantStatus.Vlan ||
		gi2011.Duplex != wantStatus.Duplex ||
		gi2011.Speed != wantStatus.Speed ||
		gi2011.Type != wantStatus.Type {
		t.Errorf("Status fail result %s, want Status %s", sw.Status()[60], wantStatus)
	}

	if err := sw.ParseInterfaceStatus(CiscoNXResponse); err != nil {
		t.Fatalf("Error parsing C2960 response: %s", err)
	}

	if len(sw.Status()) != CiscoNXInterfacesCount {
		t.Errorf("Error want %d interfaces got %d", CiscoNXInterfacesCount, len(sw.Status()))
	}
	fmt.Println(sw.ReturnIcingaResult())

}

func TestCheck_InterfaceStatus(t *testing.T) {
	_, err := sw.CheckInterfaceStatus(params.host, params.username, params.password, params.identity, params.port)
	if err != nil {
		t.Errorf("Error CheckInterfaceStatus: %s", err)
	}
}
