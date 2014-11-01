package search

import "github.com/dertseha/everoute/universe"

type SolarSystemJumpData struct {
	FromSolarSystemId universe.Id
	ToSolarSystemId   universe.Id
}

type SolarSystemData struct {
	RegionId        universe.Id
	ConstellationId universe.Id
	SolarSystemId   universe.Id
	Name            string
	X               float64
	Y               float64
	Z               float64
	Security        float64
}

var SolarSystems = []SolarSystemData{
	{10000030, 20000367, 30002505, "Hulm", -8.68416973876787E+16, 3.75027957308021E+16, 3.0503349623728E+15, 1},
	{10000030, 20000367, 30002506, "Osoggur", -1.19198464346739E+17, 4.32706453871816E+16, -183855711628938, 0.53156828951296},
	{10000030, 20000367, 30002507, "Abudban", -1.07534933500286E+17, 4.80440858050229E+16, 583532325443579, 0.73244772463105},
	{10000030, 20000367, 30002508, "Trytedald", -9.42913103547792E+16, 4.1055540147143E+16, -1.08455471398144E+16, 0.87965850369903},
	{10000030, 20000367, 30002509, "Odatrik", -9.88716571550024E+16, 4.27212367929229E+16, -1.40362956832075E+16, 0.81617854816542},
	{10000030, 20000367, 30002510, "Rens", -9.91223410909671E+16, 4.03351006314892E+16, -2.86772881958385E+15, 0.89458186592099},
	{10000030, 20000367, 30002511, "Ameinaka", -8.56875728236928E+16, 3.58909991782934E+16, 7.47763282842871E+15, 0.98742320531048},
	{10000030, 20000368, 30002512, "Alakgur", -1.16054722259902E+17, 5.09215090000822E+16, -5.16556274744165E+15, 0.55883569801702},
	{10000030, 20000368, 30002513, "Dammalin", -1.1815410366612E+17, 5.6822224460478E+16, -2.65480104773659E+15, 0.47875580238313},
	{10000030, 20000368, 30002514, "Bosboger", -1.25234161059748E+17, 6.30421343236535E+16, 1.8468832071076E+15, 0.32458180474549},
	{10000030, 20000368, 30002515, "Olfeim", -1.22706891786159E+17, 5.11951810420961E+16, 3.41678227554805E+15, 0.43868425633026},
	{10000030, 20000368, 30002516, "Lulm", -1.20046221138605E+17, 6.73406390348985E+16, 2.42128266201547E+15, 0.34397383105912},
	{10000030, 20000368, 30002517, "Gulmorogod", -1.20834247838822E+17, 6.23588997692797E+16, 7.84955408194756E+15, 0.37679403765006},
	{10000030, 20000369, 30002518, "Edmalbrurdus", -8.53932408230673E+16, 3.49885542376397E+16, -122058034903028, 0.98923024302057},
	{10000030, 20000369, 30002519, "Kronsur", -8.82833856697144E+16, 2.49992067701992E+16, -3.07164355673895E+15, 0.88224525941906},
	{10000030, 20000369, 30002520, "Dumkirinur", -8.14177504035524E+16, 1.5689258993142E+16, 1.69333239924741E+15, 0.70661563468129},
	{10000030, 20000369, 30002521, "Sist", -8.64929469258681E+16, 1.95917659941832E+16, 3.69595782955841E+15, 0.8080591710965},
	{10000030, 20000369, 30002522, "Obrolber", -7.94811388313164E+16, 1.01521731805444E+16, 6.74417034683866E+15, 0.56718225299904},
	{10000030, 20000369, 30002523, "Austraka", -7.33920094715565E+16, 2.09646204697357E+16, 1.30568415428305E+16, 0.75571051812847},
	{10000030, 20000370, 30002524, "Ivar", -9.24445035896468E+16, 4.13141065643612E+16, 9.22010826826928E+15, 0.95621091866395},
	{10000030, 20000370, 30002525, "Meirakulf", -9.46065946032824E+16, 4.28908698986709E+16, 1.72844955406648E+16, 0.88206671595594},
	{10000030, 20000370, 30002526, "Frarn", -9.63572917589253E+16, 4.07174107458614E+16, 2.44329604056643E+16, 0.83897091308227},
	{10000030, 20000370, 30002527, "Illinfrik", -9.38085331885443E+16, 4.53570405662506E+16, 2.91014806364526E+16, 0.84635978457726},
	{10000030, 20000370, 30002528, "Balginia", -9.60969703770856E+16, 5.06120508549049E+16, 3.13603691228973E+16, 0.81765668524698},
	{10000030, 20000370, 30002529, "Gyng", -1.00859177981538E+17, 4.20955406264848E+16, 2.48367693731105E+16, 0.79814297004297},
	{10000030, 20000370, 30002530, "Avesber", -1.05003036781673E+17, 3.98059970105447E+16, 2.42782886010607E+16, 0.75318365680617},
	{10000030, 20000371, 30002531, "Gerek", -7.28283439746005E+16, 2.02699364359682E+16, 2.16022311464335E+16, 0.7322521435152},
	{10000030, 20000371, 30002532, "Tongofur", -7.14155974564457E+16, 1.34693934700295E+16, 1.86392150729702E+16, 0.58306899552179},
	{10000030, 20000371, 30002533, "Gerbold", -6.72585993604215E+16, 2.56109692024147E+16, 2.98028126287425E+16, 0.82256580370087},
	{10000030, 20000371, 30002534, "Rokofur", -7.78207908882556E+16, 1.4239548450016E+16, 2.49920425984915E+16, 0.6180885585596},
	{10000030, 20000371, 30002535, "Ebasgerdur", -7.57865247722559E+16, 2.23471000752983E+16, 2.32403265220877E+16, 0.77438187323098},
	{10000030, 20000371, 30002536, "Ebodold", -7.1451550277843E+16, 7.34804289590815E+15, 1.41668729544247E+16, 0.44463154961851},
	{10000030, 20000372, 30002537, "Amamake", -1.24292266288494E+17, 4.41943641937272E+16, 6.1103924335871E+15, 0.43876124811132},
	{10000030, 20000372, 30002538, "Vard", -1.28361445920953E+17, 3.73376035618511E+16, 5.42256160129145E+15, 0.37909593843813},
	{10000030, 20000372, 30002539, "Siseide", -1.3224641056418E+17, 4.0994863783094E+16, 5.45487831030784E+15, 0.33835110274775},
	{10000030, 20000372, 30002540, "Lantorn", -1.30875230127806E+17, 3.22310739458827E+16, 8.11293222582307E+15, 0.33948597391556},
	{10000030, 20000372, 30002541, "Dal", -1.25768310960472E+17, 3.44540882367103E+16, 9.68965601729805E+15, 0.41135958898872},
	{10000030, 20000372, 30002542, "Auga", -1.28081842291416E+17, 4.27330971908584E+16, 6.56850905029077E+15, 0.38548382057593},
	{10000030, 20000373, 30002543, "Eystur", -9.82937414808922E+16, 4.8222147988334E+16, 6.4110148903271E+16, 0.94923182184583},
	{10000030, 20000373, 30002544, "Pator", -9.35790295622701E+16, 4.37806703802814E+16, 5.6561954432336E+16, 1},
	{10000030, 20000373, 30002545, "Lustrevik", -9.71092920690204E+16, 3.42746659215896E+16, 5.48217076517325E+16, 0.94626341602458},
	{10000030, 20000373, 30002546, "Isendeldik", -9.03945968562323E+16, 2.15126755517656E+16, 4.90711432249191E+16, 0.80885548037394},
	{10000030, 20000373, 30002547, "Ammold", -8.53486657983356E+16, 1.87491150576755E+16, 4.54301567684968E+16, 1},
	{10000030, 20000373, 30002548, "Emolgranlan", -8.83171240857611E+16, 1.03068549565459E+16, 3.90433995489847E+16, 0.5302433793008},
	{10000030, 20000374, 30002549, "Offugen", -6.81987799905678E+16, 1.73224771519956E+16, 3.28420086627681E+16, 0.6498344032544},
	{10000030, 20000374, 30002550, "Roniko", -6.10636162718802E+16, 1.6690065913918E+16, 3.1437935845438E+16, 0.59364460351451},
	{10000030, 20000374, 30002551, "Aralgrund", -5.91847410905733E+16, 5.94327316342335E+15, 2.61675700162955E+16, 0.32155892279011},
	{10000030, 20000374, 30002552, "Eddar", -6.88716576088479E+16, 1.50293388056657E+16, 3.8079835239106E+16, 0.58706356658978},
	{10000030, 20000374, 30002553, "Bogelek", -7.03642772158002E+16, 4.40178950075128E+15, 2.46135852625354E+16, 0.3663701066351},
	{10000030, 20000374, 30002554, "Wiskeber", -5.56922522203614E+16, 1.14043942026213E+16, 2.51855786896233E+16, 0.40443614610229},
	{10000030, 20000375, 30002555, "Eifer", -8.51465442882317E+16, 6.55251738743346E+15, 3.83703343316904E+16, 0.44452511601023},
	{10000030, 20000375, 30002556, "Gusandall", -7.86394955637697E+16, 3.96887663067516E+15, 4.14450877577781E+16, 0.3765845900301},
	{10000030, 20000375, 30002557, "Atgur", -8.2094100444973E+16, 8.40487487971838E+15, 3.30049440251471E+16, 0.4830198834755},
	{10000030, 20000375, 30002558, "Endrulf", -7.89909386248656E+16, 1.46548284082869E+16, 3.3847481147147E+16, 0.61461149570284},
	{10000030, 20000375, 30002559, "Ingunn", -8.40318449687837E+16, -2.0392984029254E+15, 3.4774824111634E+16, 0.27679566348469},
	{10000030, 20000375, 30002560, "Gultratren", -8.89941583123118E+16, -450001015918553, 2.6557255711959E+16, 0.30927565721679},
	{10000030, 20000375, 30002561, "Auren", -7.33544896455761E+16, 7.03781448377451E+15, 4.23567533201794E+16, 0.41791273595739},
	{10000030, 20000376, 30002562, "Trer", -6.323603259434E+16, 2.19876104828077E+16, 3.95674411670798E+16, 0.71266511887127},
	{10000030, 20000376, 30002563, "Egmur", -6.36566592720724E+16, 2.09565003701098E+16, 4.4901627485029E+16, 0.65550527986505},
	{10000030, 20000376, 30002564, "Javrendei", -5.72955737063684E+16, 3.33810566607024E+16, 4.03918731338839E+16, 0.87865238571603},
	{10000030, 20000376, 30002565, "Appen", -5.71784272452761E+16, 3.45326630343797E+16, 4.57445618756747E+16, 0.82132929535884},
	{10000030, 20000376, 30002566, "Klir", -5.18240170058142E+16, 3.42935657881092E+16, 4.50954541707727E+16, 0.75694238370415},
	{10000030, 20000376, 30002567, "Jorus", -5.15218058416557E+16, 3.96878808456501E+16, 4.89308337890348E+16, 0.72698042807852},
	{10000030, 20000377, 30002568, "Onga", -9.81034571079606E+16, 4.25572692025824E+16, 4.76921719068648E+16, 0.95351375759448},
	{10000030, 20000377, 30002569, "Osaumuni", -9.44907981202691E+16, 4.15212960098381E+16, 4.15140911322729E+16, 0.91486754052731},
	{10000030, 20000377, 30002570, "Magiko", -9.45651315167908E+16, 3.58211325126733E+16, 4.78905708617062E+16, 0.93830202528173},
	{10000030, 20000377, 30002571, "Oremmulf", -9.2989865297084E+16, 4.59504498656612E+16, 3.89206067221998E+16, 0.89710667044381},
	{10000030, 20000377, 30002572, "Hurjafren", -9.04668261998736E+16, 5.20494372315183E+16, 3.91374308543389E+16, 0.87692885085938},
	{10000030, 20000377, 30002573, "Vullat", -8.89857715276627E+16, 4.17494051023919E+16, 5.00580360386038E+16, 0.97136261120209},
	{10000030, 20000378, 30002574, "Hrondedir", -6.49191352082384E+16, 4.08762334606926E+15, 3.72759766437149E+16, 0.31742355818347},
	{10000030, 20000378, 30002575, "Sotrenzur", -7.22048527469629E+16, -188500197773560, 3.66263134980257E+16, 0.27596997895457},
	{10000030, 20000378, 30002576, "Hrondmund", -5.90047783483382E+16, 8.80927657676082E+15, 4.18344994116077E+16, 0.35880389974293},
	{10000030, 20000378, 30002577, "Bundindus", -5.97133365855849E+16, 1.30933971375194E+16, 4.75231156629862E+16, 0.428685119678},
	{10000030, 20000378, 30002578, "Otraren", -5.82600134791362E+16, 1.76470259196202E+16, 5.14625208380162E+16, 0.47510923142446},
	{10000030, 20000378, 30002579, "Hedgiviter", -5.24522349956376E+16, 1.68138338091099E+16, 4.80996813590176E+16, 0.41567334445711},
	{10000030, 20000378, 30002580, "Katugumur", -7.35188485541418E+16, -7.41160744624274E+15, 3.2350980813408E+16, 0.17129139817121},
	{10000030, 20000367, 30012505, "Malukker", -8.26875728236928E+16, 3.58909991782934E+16, 4.47763282842871E+15, 0.98801167454022},
	{10000030, 20000373, 30012547, "Hadaugago", -1.01293741480892E+17, 4.8222147988334E+16, 6.1110148903271E+16, 0.94877196464412},
	{10000030, 20000377, 30022547, "Krilmokenur", -1.01103457107961E+17, 4.25572692025824E+16, 4.46921719068648E+16, 0.91170062762363},
	{10000030, 20000370, 30042505, "Usteli", -8.94445035896468E+16, 4.13141065643612E+16, 1.22201082682693E+16, 0.95059503262015},
	{10000030, 20000377, 30042547, "Loguttur", -8.99857715276627E+16, 4.17494051023919E+16, 4.70580360386038E+16, 0.95733121058868},
	{10000030, 20000369, 30032505, "Todeko", -8.43932408230673E+16, 3.49885542376397E+16, 1.87794196509697E+15, 0.99211814063388},
	{10000030, 20000377, 30032547, "Larkugei", -9.34907981202691E+16, 4.15212960098381E+16, 4.35140911322729E+16, 0.93264311721317}}

var SolarSystemJumps = []SolarSystemJumpData{
	{30002505, 30002511},
	{30002505, 30002518},
	{30002506, 30002507},
	{30002506, 30002537},
	{30002507, 30002509},
	{30002507, 30002510},
	{30002507, 30002512},
	{30002507, 30002530},
	{30002508, 30002509},
	{30002508, 30002524},
	{30002509, 30002510},
	{30002510, 30002526},
	{30002511, 30002524},
	{30002511, 30012505},
	{30002512, 30002513},
	{30002513, 30002514},
	{30002513, 30002515},
	{30002514, 30002516},
	{30002514, 30002517},
	{30002516, 30002517},
	{30002517, 30002537},
	{30002518, 30002519},
	{30002518, 30032505},
	{30002519, 30002520},
	{30002519, 30002521},
	{30002520, 30002521},
	{30002520, 30002522},
	{30002522, 30002523},
	{30002523, 30002531},
	{30002524, 30002525},
	{30002524, 30042505},
	{30002525, 30002526},
	{30002526, 30002527},
	{30002526, 30002529},
	{30002526, 30002530},
	{30002527, 30002528},
	{30002527, 30002569},
	{30002528, 30002572},
	{30002529, 30002530},
	{30002529, 30002568},
	{30002531, 30002532},
	{30002531, 30002533},
	{30002531, 30002534},
	{30002531, 30002535},
	{30002532, 30002536},
	{30002533, 30002549},
	{30002534, 30002535},
	{30002534, 30002560},
	{30002537, 30002538},
	{30002537, 30002539},
	{30002537, 30002541},
	{30002537, 30002542},
	{30002538, 30002539},
	{30002538, 30002540},
	{30002538, 30002541},
	{30002539, 30002540},
	{30002539, 30002541},
	{30002539, 30002542},
	{30002540, 30002541},
	{30002541, 30002542},
	{30002543, 30002544},
	{30002543, 30002545},
	{30002543, 30002573},
	{30002543, 30012547},
	{30002544, 30002545},
	{30002544, 30002547},
	{30002544, 30002568},
	{30002545, 30002546},
	{30002545, 30002568},
	{30002546, 30002547},
	{30002547, 30002548},
	{30002548, 30002555},
	{30002549, 30002550},
	{30002549, 30002552},
	{30002550, 30002551},
	{30002550, 30002574},
	{30002551, 30002553},
	{30002551, 30002554},
	{30002552, 30002561},
	{30002552, 30002562},
	{30002553, 30002580},
	{30002555, 30002556},
	{30002555, 30002557},
	{30002555, 30002558},
	{30002556, 30002559},
	{30002556, 30002561},
	{30002556, 30002574},
	{30002556, 30002575},
	{30002557, 30002558},
	{30002559, 30002560},
	{30002559, 30002575},
	{30002562, 30002563},
	{30002562, 30002564},
	{30002564, 30002565},
	{30002565, 30002566},
	{30002566, 30002567},
	{30002568, 30002569},
	{30002568, 30002570},
	{30002568, 30022547},
	{30002569, 30002571},
	{30002569, 30032547},
	{30002570, 30002573},
	{30002571, 30002572},
	{30002573, 30042547},
	{30002574, 30002575},
	{30002574, 30002576},
	{30002575, 30002580},
	{30002576, 30002577},
	{30002577, 30002578},
	{30002578, 30002579}}
