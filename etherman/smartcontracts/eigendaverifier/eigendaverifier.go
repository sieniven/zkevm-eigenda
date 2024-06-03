// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eigendaverifier

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// EigenDARollupUtilsBlobVerificationProof is an auto generated low-level Go binding around an user-defined struct.
type EigenDARollupUtilsBlobVerificationProof struct {
	BatchId        uint32
	BlobIndex      uint32
	BatchMetadata  IEigenDAServiceManagerBatchMetadata
	InclusionProof []byte
	QuorumIndices  []byte
}

// EigenDAVerifierBlobData is an auto generated low-level Go binding around an user-defined struct.
type EigenDAVerifierBlobData struct {
	BlobHeader            IEigenDAServiceManagerBlobHeader
	BlobVerificationProof EigenDARollupUtilsBlobVerificationProof
}

// IEigenDAServiceManagerBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBatchHeader struct {
	BlobHeadersRoot       [32]byte
	QuorumNumbers         []byte
	SignedStakeForQuorums []byte
	ReferenceBlockNumber  uint32
}

// IEigenDAServiceManagerBatchMetadata is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBatchMetadata struct {
	BatchHeader             IEigenDAServiceManagerBatchHeader
	SignatoryRecordHash     [32]byte
	ConfirmationBlockNumber uint32
}

// IEigenDAServiceManagerBlobHeader is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBlobHeader struct {
	Commitment       BN254G1Point
	DataLength       uint32
	QuorumBlobParams []IEigenDAServiceManagerQuorumBlobParam
}

// IEigenDAServiceManagerQuorumBlobParam is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerQuorumBlobParam struct {
	QuorumNumber                    uint8
	AdversaryThresholdPercentage    uint8
	ConfirmationThresholdPercentage uint8
	ChunkLength                     uint32
}

// EigendaverifierMetaData contains all meta data concerning the Eigendaverifier contract.
var EigendaverifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eigenDAServiceManager\",\"type\":\"address\",\"internalType\":\"contractIEigenDAServiceManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptAdminRole\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decodeBlobData\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"blobData\",\"type\":\"tuple\",\"internalType\":\"structEigenDAVerifier.BlobData\",\"components\":[{\"name\":\"blobHeader\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BlobHeader\",\"components\":[{\"name\":\"commitment\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"dataLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBlobParams\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenDAServiceManager.QuorumBlobParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"adversaryThresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"confirmationThresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chunkLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]},{\"name\":\"blobVerificationProof\",\"type\":\"tuple\",\"internalType\":\"structEigenDARollupUtils.BlobVerificationProof\",\"components\":[{\"name\":\"batchId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blobIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"batchMetadata\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BatchMetadata\",\"components\":[{\"name\":\"batchHeader\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BatchHeader\",\"components\":[{\"name\":\"blobHeadersRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signedStakeForQuorums\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"confirmationBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"inclusionProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumIndices\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getProcotolName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"pendingAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDataAvailabilityProtocol\",\"inputs\":[{\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\",\"internalType\":\"contractIEigenDAServiceManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferAdminRole\",\"inputs\":[{\"name\":\"newPendingAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AcceptAdminRole\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDataAvailabilityProtocol\",\"inputs\":[{\"name\":\"newTrustedSequencer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIEigenDAServiceManager\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferAdminRole\",\"inputs\":[{\"name\":\"newPendingAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BatchAlreadyVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedMaxVerifyBatches\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FinalPendingStateNumInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchTimeoutNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchesAlreadyActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchesDecentralized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchesOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForcedDataDoesNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalExitRootNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HaltTimeoutNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HugeTokenMetadataNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InitSequencedBatchDoesNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitializeTransaction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRangeBatchTimeTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRangeForceBatchTimeout\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRangeMultiplierBatchFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxTimestampSequenceInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewAccInputHashDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewPendingStateTimeoutMustBeLower\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewStateRootNotInsidePrime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughMaticAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughPOLAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OldAccInputHashDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OldStateRootDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPendingAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRollupManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTrustedAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTrustedSequencer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingStateDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingStateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingStateNotConsolidable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequenceZeroBatches\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequencedTimestampInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransactionsLengthAboveMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"inputs\":[]}]",
	Bin: "0x363038303630343035323334383031353631303031303537363030303830666435623530363034303531363130646330333830333830363130646330383333393831303136303430383139303532363130303266393136313030373835363562363030313830353436303031363030313630613031623033393338343136363030313630303136306130316230333139393138323136313739303931353536303030383035343932393039333136393131363137393035353631303062323536356236303031363030313630613031623033383131363831313436313030373535373630303038306664356235303536356236303030383036303430383338353033313231353631303038623537363030303830666435623832353136313030393638313631303036303536356236303230383430313531393039323530363130306137383136313030363035363562383039313530353039323530393239303530353635623631306366663830363130306331363030303339363030306633666536303830363034303532333438303135363130303130353736303030383066643562353036303034333631303631303038383537363030303335363065303163383036336162613463383064313136313030356235373830363361626134633830643134363130306564353738303633616461386639313931343631303130643537383036336534663137313230313436313031323035373830363366383531613434303134363130313439353736303030383066643562383036333236373832323437313436313030386435373830363333623531626534623134363130306264353738303633376364373662386231343631303064323537383036333863336437333031313436313030653535373562363030303830666435623630303235343631303061303930363030313630303136306130316230333136383135363562363034303531363030313630303136306130316230333930393131363831353236303230303135623630343035313830393130333930663335623631303064303631303063623336363030343631303436383536356236313031356335363562303035623631303064303631303065303336363030343631303463633536356236313031656535363562363130306430363130323665353635623631303130303631303066623336363030343631303465393536356236313032656635363562363034303531363130306234393139303631303634343536356236313030643036313031316233363630303436313034636335363562363130333061353635623630343038303531383038323031383235323630303738313532363634353639363736353665343434313630633831623630323038323031353239303531363130306234393139303631303732313536356236303031353436313030613039303630303136303031363061303162303331363831353635623630303036313031363838333833363130326566353635623830353136303030353436303230383330313531363034303531363332313934363065313630653231623831353239333934353037335f5f2433393966396365386464333361303664313434636531656232346438343565323830245f5f393336333836353138333834393336313031623839333930393236303031363030313630613031623033393039313136393136303034303136313037333435363562363030303630343035313830383330333831383638303362313538303135363130316430353736303030383066643562353035616634313538303135363130316534353733643630303038303365336436303030666435623530353035303530353035303530353035363562363030313534363030313630303136306130316230333136333331343631303231393537363034303531363334373535363537393630653031623831353236303034303136303430353138303931303339306664356236303030383035343630303136303031363061303162303331393136363030313630303136306130316230333833313639303831313739303931353536303430353139303831353237666433333162643463346364316166656362393461323235313834626465643136316666333231333632346261346662353863346633306335613836313134346139303630323030313562363034303531383039313033393061313530353635623630303235343630303136303031363061303162303331363333313436313032393935373630343035313633643165633462323336306530316238313532363030343031363034303531383039313033393066643562363030323534363030313830353436303031363030313630613031623033313931363630303136303031363061303162303339303932313639313832313739303535363034303531393038313532376630353664633438376262663037393564306262623162346630616635323361383535353033636666373430626662346435343735663761393063303931653865393036303230303136303430353138303931303339306131353635623631303266373631303338333536356236313033303338323834303138343631306230623536356239333932353035303530353635623630303135343630303136303031363061303162303331363333313436313033333535373630343035313633343735353635373936306530316238313532363030343031363034303531383039313033393066643562363030323830353436303031363030313630613031623033313931363630303136303031363061303162303338333136393038313137393039313535363034303531393038313532376661356235366237393036666430613230653366333531323064643833343364623165313265303337613663393031313163376534323838356538326131636536393036303230303136313032363335363562363034303830353136306530383130313832353236303030363061303832303138313831353236306330383330313832393035323932383230313932383335323630363038303833303139313930393135323630383038323031353239303831353236303230383130313631303431613630343038303531363061303830383230313833353236303030383038333532363032303830383430313832393035323834353136306530383130313836353236303630383038323031383438313532363038303833303138323930353239343832303135323630633038313031383339303532393238333532383230313831393035323831383430313532393039313832303139303831353236303230303136303630383135323630323030313630363038313532353039303536356239303532393035363562363030303830383336303166383430313132363130343331353736303030383066643562353038313335363766666666666666666666666666666666383131313135363130343439353736303030383066643562363032303833303139313530383336303230383238353031303131313135363130343631353736303030383066643562393235303932393035303536356236303030383036303030363034303834383630333132313536313034376435373630303038306664356238333335393235303630323038343031333536376666666666666666666666666666666638313131313536313034396235373630303038306664356236313034613738363832383730313631303431663536356239343937393039363530393339343530353035303530353635623630303136303031363061303162303338313136383131343631303463393537363030303830666435623530353635623630303036303230383238343033313231353631303464653537363030303830666435623831333536313033303338313631303462343536356236303030383036303230383338353033313231353631303466633537363030303830666435623832333536376666666666666666666666666666666638313131313536313035313335373630303038306664356236313035316638353832383630313631303431663536356239303936393039353530393335303530353035303536356236303030383135313830383435323630303035623831383131303135363130353531353736303230383138353031383130313531383638333031383230313532303136313035333535363562353036303030363032303832383630313031353236303230363031663139363031663833303131363835303130313931353035303932393135303530353635623630303036336666666666666666383038333531313638343532383036303230383430313531313636303230383530313532363034303833303135313630613036303430383630313532383035313630363036306130383730313532383035313631303130303837303135323630323038313031353136303830363130313230383830313532363130356332363130313830383830313832363130353262353635623930353036303430383230313531363066663139383838333033303136313031343038393031353236313035653038323832363130353262353635623931353035303833363036303833303135313136363130313630383830313532363032303833303135313630633038383031353238333630343038343031353131363630653038383031353236303630383630313531393335303836383130333630363038383031353236313036316638313835363130353262353635623933353035303530353036303830383330313531383438323033363038303836303135323631303633623832383236313035326235363562393539343530353035303530353035363562363030303630323038303833353238333531363034303832383530313532363065303834303136313036366636303630383630313833353138303531383235323630323039303831303135313931303135323536356238313833303135313633666666666666666631363630613038363031353236303430393039313031353136303830363063303836303138313930353238313531393238333930353239303833303139313630303039313930363130313030383730313930356238303834313031353631303666393537363130366535383238363531363066663831353131363832353236306666363032303832303135313136363032303833303135323630666636303430383230313531313636303430383330313532363366666666666666663630363038323031353131363630363038333031353235303530353635623933383530313933363030313933393039333031393239303832303139303631303661323536356235303933383730313531383638353033363031663139303136303430383830313532393336313037313538313836363130353731353635623938393735303530353035303530353035303530353635623630323038313532363030303631303330333630323038333031383436313035326235363562363036303831353236303030363065303832303136313037353636303630383430313837353138303531383235323630323039303831303135313931303135323536356236303230383638313031353136336666666666666666313636306130383530313532363034303837303135313630383036306330383630313831393035323831353139333834393035323930383230313932363030303931393036313031303038373031393035623830383431303135363130376531353736313037636438323837353136306666383135313136383235323630666636303230383230313531313636303230383330313532363066663630343038323031353131363630343038333031353236336666666666666666363036303832303135313136363036303833303135323530353035363562393438343031393436303031393339303933303139323930383230313930363130373861353635623530363030313630303136306130316230333839313638373835303135323836383130333630343038383031353236313038303238313839363130353731353635623961393935303530353035303530353035303530353035303536356236333465343837623731363065303162363030303532363034313630303435323630323436303030666435623630343035313630363038313031363766666666666666666666666666666666383131313832383231303137313536313038343935373631303834393631303831303536356236303430353239303536356236303430353136303830383130313637666666666666666666666666666666663831313138323832313031373135363130383439353736313038343936313038313035363562363034303531363061303831303136376666666666666666666666666666666638313131383238323130313731353631303834393537363130383439363130383130353635623630343038303531393038313031363766666666666666666666666666666666383131313832383231303137313536313038343935373631303834393631303831303536356236303430353136303166383230313630316631393136383130313637666666666666666666666666666666663831313138323832313031373135363130386531353736313038653136313038313035363562363034303532393139303530353635623830333536336666666666666666383131363831313436313038666435373630303038306664356239313930353035363562383033353630666638313136383131343631303866643537363030303830666435623630303038323630316638333031313236313039323435373630303038306664356238313335363766666666666666666666666666666666383131313135363130393365353736313039336536313038313035363562363130393531363031663832303136303166313931363630323030313631303862383536356238313831353238343630323038333836303130313131313536313039363635373630303038306664356238313630323038353031363032303833303133373630303039313831303136303230303139313930393135323933393235303530353035363562363030303630363038323834303331323135363130393935353736303030383066643562363130393964363130383236353635623930353038313335363766666666666666666666666666666666383038323131313536313039623735373630303038306664356239303833303139303630383038323836303331323135363130396362353736303030383066643562363130396433363130383466353635623832333538313532363032303833303133353832383131313135363130396539353736303030383066643562363130396635383738323836303136313039313335363562363032303833303135323530363034303833303133353832383131313135363130613064353736303030383066643562363130613139383738323836303136313039313335363562363034303833303135323530363130613262363036303834303136313038653935363562363036303832303135323833353235303530363032303832383130313335393038323031353236313061346136303430383330313631303865393536356236303430383230313532393239313530353035363562363030303630613038323834303331323135363130613637353736303030383066643562363130613666363130383732353635623930353036313061376138323631303865393536356238313532363130613838363032303833303136313038653935363562363032303832303135323630343038323031333536376666666666666666666666666666666638303832313131353631306161383537363030303830666435623631306162343835383338363031363130393833353635623630343038343031353236303630383430313335393135303830383231313135363130616364353736303030383066643562363130616439383538333836303136313039313335363562363036303834303135323630383038343031333539313530383038323131313536313061663235373630303038306664356235303631306166663834383238353031363130393133353635623630383038333031353235303932393135303530353635623630303036303230383038333835303331323135363130623165353736303030383066643562383233353637666666666666666666666666666666663830383231313135363130623336353736303030383066643562383138353031393135303630343038303833383830333132313536313062346335373630303038306664356236313062353436313038393535363562383333353833383131313135363130623633353736303030383066643562383430313830383930333630383038313132313536313062373635373630303038306664356236313062376536313038323635363562383438323132313536313062386235373630303038306664356236313062393336313038393535363562393135303832333538323532383738333031333538383833303135323831383135323631306261663835383430313631303865393536356238383832303135323630363039313530383138333031333538363831313131353631306263383537363030303830666435623830383430313933353035303861363031663834303131323631306264643537363030303830666435623832333538363831313131353631306265663537363130626566363130383130353635623631306266643839383236303035316230313631303862383536356238313831353236303037393139303931316238343031383930313930383938313031393038643833313131353631306331633537363030303830666435623934386130313934356238323836313031353631306338633537363038303836386630333132313536313063336135373630303038303831666435623631306334323631303834663536356236313063346238373631303930323536356238313532363130633538386338383031363130393032353635623863383230313532363130633637383938383031363130393032353635623839383230313532363130633736383638383031363130386539353635623831383730313532383235323630383039353930393530313934393038613031393036313063323135363562393638333031393639303936353235303833353235303530383338353031333539313530383238323131313536313063616235373630303038306664356236313063623738383833383630313631306135353536356238353832303135323830393535303530353035303530353039323931353035303536666561323634363937303636373335383232313232303739656631306465393939356166306232336264626338393166333561373761643930633831346539666535623937623439626264626235663530306366306236343733366636633633343330303038313430303333",
}

// EigendaverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use EigendaverifierMetaData.ABI instead.
var EigendaverifierABI = EigendaverifierMetaData.ABI

// EigendaverifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EigendaverifierMetaData.Bin instead.
var EigendaverifierBin = EigendaverifierMetaData.Bin

// DeployEigendaverifier deploys a new Ethereum contract, binding an instance of Eigendaverifier to it.
func DeployEigendaverifier(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _eigenDAServiceManager common.Address) (common.Address, *types.Transaction, *Eigendaverifier, error) {
	parsed, err := EigendaverifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EigendaverifierBin), backend, _admin, _eigenDAServiceManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eigendaverifier{EigendaverifierCaller: EigendaverifierCaller{contract: contract}, EigendaverifierTransactor: EigendaverifierTransactor{contract: contract}, EigendaverifierFilterer: EigendaverifierFilterer{contract: contract}}, nil
}

// Eigendaverifier is an auto generated Go binding around an Ethereum contract.
type Eigendaverifier struct {
	EigendaverifierCaller     // Read-only binding to the contract
	EigendaverifierTransactor // Write-only binding to the contract
	EigendaverifierFilterer   // Log filterer for contract events
}

// EigendaverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigendaverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigendaverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigendaverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigendaverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigendaverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigendaverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigendaverifierSession struct {
	Contract     *Eigendaverifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EigendaverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigendaverifierCallerSession struct {
	Contract *EigendaverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EigendaverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigendaverifierTransactorSession struct {
	Contract     *EigendaverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EigendaverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigendaverifierRaw struct {
	Contract *Eigendaverifier // Generic contract binding to access the raw methods on
}

// EigendaverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigendaverifierCallerRaw struct {
	Contract *EigendaverifierCaller // Generic read-only contract binding to access the raw methods on
}

// EigendaverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigendaverifierTransactorRaw struct {
	Contract *EigendaverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigendaverifier creates a new instance of Eigendaverifier, bound to a specific deployed contract.
func NewEigendaverifier(address common.Address, backend bind.ContractBackend) (*Eigendaverifier, error) {
	contract, err := bindEigendaverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eigendaverifier{EigendaverifierCaller: EigendaverifierCaller{contract: contract}, EigendaverifierTransactor: EigendaverifierTransactor{contract: contract}, EigendaverifierFilterer: EigendaverifierFilterer{contract: contract}}, nil
}

// NewEigendaverifierCaller creates a new read-only instance of Eigendaverifier, bound to a specific deployed contract.
func NewEigendaverifierCaller(address common.Address, caller bind.ContractCaller) (*EigendaverifierCaller, error) {
	contract, err := bindEigendaverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigendaverifierCaller{contract: contract}, nil
}

// NewEigendaverifierTransactor creates a new write-only instance of Eigendaverifier, bound to a specific deployed contract.
func NewEigendaverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*EigendaverifierTransactor, error) {
	contract, err := bindEigendaverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigendaverifierTransactor{contract: contract}, nil
}

// NewEigendaverifierFilterer creates a new log filterer instance of Eigendaverifier, bound to a specific deployed contract.
func NewEigendaverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*EigendaverifierFilterer, error) {
	contract, err := bindEigendaverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigendaverifierFilterer{contract: contract}, nil
}

// bindEigendaverifier binds a generic wrapper to an already deployed contract.
func bindEigendaverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigendaverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eigendaverifier *EigendaverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eigendaverifier.Contract.EigendaverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eigendaverifier *EigendaverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.EigendaverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eigendaverifier *EigendaverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.EigendaverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eigendaverifier *EigendaverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eigendaverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eigendaverifier *EigendaverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eigendaverifier *EigendaverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Eigendaverifier *EigendaverifierCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Eigendaverifier *EigendaverifierSession) Admin() (common.Address, error) {
	return _Eigendaverifier.Contract.Admin(&_Eigendaverifier.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Eigendaverifier *EigendaverifierCallerSession) Admin() (common.Address, error) {
	return _Eigendaverifier.Contract.Admin(&_Eigendaverifier.CallOpts)
}

// DecodeBlobData is a free data retrieval call binding the contract method 0xaba4c80d.
//
// Solidity: function decodeBlobData(bytes data) pure returns((((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[]),(uint32,uint32,((bytes32,bytes,bytes,uint32),bytes32,uint32),bytes,bytes)) blobData)
func (_Eigendaverifier *EigendaverifierCaller) DecodeBlobData(opts *bind.CallOpts, data []byte) (EigenDAVerifierBlobData, error) {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "decodeBlobData", data)

	if err != nil {
		return *new(EigenDAVerifierBlobData), err
	}

	out0 := *abi.ConvertType(out[0], new(EigenDAVerifierBlobData)).(*EigenDAVerifierBlobData)

	return out0, err

}

// DecodeBlobData is a free data retrieval call binding the contract method 0xaba4c80d.
//
// Solidity: function decodeBlobData(bytes data) pure returns((((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[]),(uint32,uint32,((bytes32,bytes,bytes,uint32),bytes32,uint32),bytes,bytes)) blobData)
func (_Eigendaverifier *EigendaverifierSession) DecodeBlobData(data []byte) (EigenDAVerifierBlobData, error) {
	return _Eigendaverifier.Contract.DecodeBlobData(&_Eigendaverifier.CallOpts, data)
}

// DecodeBlobData is a free data retrieval call binding the contract method 0xaba4c80d.
//
// Solidity: function decodeBlobData(bytes data) pure returns((((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[]),(uint32,uint32,((bytes32,bytes,bytes,uint32),bytes32,uint32),bytes,bytes)) blobData)
func (_Eigendaverifier *EigendaverifierCallerSession) DecodeBlobData(data []byte) (EigenDAVerifierBlobData, error) {
	return _Eigendaverifier.Contract.DecodeBlobData(&_Eigendaverifier.CallOpts, data)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Eigendaverifier *EigendaverifierCaller) GetProcotolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "getProcotolName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Eigendaverifier *EigendaverifierSession) GetProcotolName() (string, error) {
	return _Eigendaverifier.Contract.GetProcotolName(&_Eigendaverifier.CallOpts)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Eigendaverifier *EigendaverifierCallerSession) GetProcotolName() (string, error) {
	return _Eigendaverifier.Contract.GetProcotolName(&_Eigendaverifier.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Eigendaverifier *EigendaverifierCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Eigendaverifier *EigendaverifierSession) PendingAdmin() (common.Address, error) {
	return _Eigendaverifier.Contract.PendingAdmin(&_Eigendaverifier.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Eigendaverifier *EigendaverifierCallerSession) PendingAdmin() (common.Address, error) {
	return _Eigendaverifier.Contract.PendingAdmin(&_Eigendaverifier.CallOpts)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes data) view returns()
func (_Eigendaverifier *EigendaverifierCaller) VerifyMessage(opts *bind.CallOpts, arg0 [32]byte, data []byte) error {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "verifyMessage", arg0, data)

	if err != nil {
		return err
	}

	return err

}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes data) view returns()
func (_Eigendaverifier *EigendaverifierSession) VerifyMessage(arg0 [32]byte, data []byte) error {
	return _Eigendaverifier.Contract.VerifyMessage(&_Eigendaverifier.CallOpts, arg0, data)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes data) view returns()
func (_Eigendaverifier *EigendaverifierCallerSession) VerifyMessage(arg0 [32]byte, data []byte) error {
	return _Eigendaverifier.Contract.VerifyMessage(&_Eigendaverifier.CallOpts, arg0, data)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Eigendaverifier *EigendaverifierTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eigendaverifier.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Eigendaverifier *EigendaverifierSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Eigendaverifier.Contract.AcceptAdminRole(&_Eigendaverifier.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Eigendaverifier *EigendaverifierTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Eigendaverifier.Contract.AcceptAdminRole(&_Eigendaverifier.TransactOpts)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Eigendaverifier *EigendaverifierTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Eigendaverifier *EigendaverifierSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.SetDataAvailabilityProtocol(&_Eigendaverifier.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Eigendaverifier *EigendaverifierTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.SetDataAvailabilityProtocol(&_Eigendaverifier.TransactOpts, newDataAvailabilityProtocol)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Eigendaverifier *EigendaverifierTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Eigendaverifier *EigendaverifierSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.TransferAdminRole(&_Eigendaverifier.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Eigendaverifier *EigendaverifierTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.TransferAdminRole(&_Eigendaverifier.TransactOpts, newPendingAdmin)
}

// EigendaverifierAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Eigendaverifier contract.
type EigendaverifierAcceptAdminRoleIterator struct {
	Event *EigendaverifierAcceptAdminRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EigendaverifierAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigendaverifierAcceptAdminRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EigendaverifierAcceptAdminRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EigendaverifierAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigendaverifierAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigendaverifierAcceptAdminRole represents a AcceptAdminRole event raised by the Eigendaverifier contract.
type EigendaverifierAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*EigendaverifierAcceptAdminRoleIterator, error) {

	logs, sub, err := _Eigendaverifier.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &EigendaverifierAcceptAdminRoleIterator{contract: _Eigendaverifier.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *EigendaverifierAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Eigendaverifier.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigendaverifierAcceptAdminRole)
				if err := _Eigendaverifier.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) ParseAcceptAdminRole(log types.Log) (*EigendaverifierAcceptAdminRole, error) {
	event := new(EigendaverifierAcceptAdminRole)
	if err := _Eigendaverifier.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigendaverifierSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the Eigendaverifier contract.
type EigendaverifierSetDataAvailabilityProtocolIterator struct {
	Event *EigendaverifierSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EigendaverifierSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigendaverifierSetDataAvailabilityProtocol)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EigendaverifierSetDataAvailabilityProtocol)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EigendaverifierSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigendaverifierSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigendaverifierSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the Eigendaverifier contract.
type EigendaverifierSetDataAvailabilityProtocol struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newTrustedSequencer)
func (_Eigendaverifier *EigendaverifierFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*EigendaverifierSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _Eigendaverifier.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &EigendaverifierSetDataAvailabilityProtocolIterator{contract: _Eigendaverifier.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newTrustedSequencer)
func (_Eigendaverifier *EigendaverifierFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *EigendaverifierSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _Eigendaverifier.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigendaverifierSetDataAvailabilityProtocol)
				if err := _Eigendaverifier.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDataAvailabilityProtocol is a log parse operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newTrustedSequencer)
func (_Eigendaverifier *EigendaverifierFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*EigendaverifierSetDataAvailabilityProtocol, error) {
	event := new(EigendaverifierSetDataAvailabilityProtocol)
	if err := _Eigendaverifier.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigendaverifierTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Eigendaverifier contract.
type EigendaverifierTransferAdminRoleIterator struct {
	Event *EigendaverifierTransferAdminRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EigendaverifierTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigendaverifierTransferAdminRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EigendaverifierTransferAdminRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EigendaverifierTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigendaverifierTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigendaverifierTransferAdminRole represents a TransferAdminRole event raised by the Eigendaverifier contract.
type EigendaverifierTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*EigendaverifierTransferAdminRoleIterator, error) {

	logs, sub, err := _Eigendaverifier.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &EigendaverifierTransferAdminRoleIterator{contract: _Eigendaverifier.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *EigendaverifierTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Eigendaverifier.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigendaverifierTransferAdminRole)
				if err := _Eigendaverifier.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) ParseTransferAdminRole(log types.Log) (*EigendaverifierTransferAdminRole, error) {
	event := new(EigendaverifierTransferAdminRole)
	if err := _Eigendaverifier.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
