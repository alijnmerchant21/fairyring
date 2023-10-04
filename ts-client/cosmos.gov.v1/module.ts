// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgVoteWeighted } from "./types/cosmos/gov/v1/tx";
import { MsgSubmitProposal } from "./types/cosmos/gov/v1/tx";
import { MsgDeposit } from "./types/cosmos/gov/v1/tx";
import { MsgUpdateParams } from "./types/cosmos/gov/v1/tx";
import { MsgVote } from "./types/cosmos/gov/v1/tx";

import { WeightedVoteOption as typeWeightedVoteOption} from "./types"
import { Deposit as typeDeposit} from "./types"
import { Proposal as typeProposal} from "./types"
import { TallyResult as typeTallyResult} from "./types"
import { Vote as typeVote} from "./types"
import { DepositParams as typeDepositParams} from "./types"
import { VotingParams as typeVotingParams} from "./types"
import { TallyParams as typeTallyParams} from "./types"
import { Params as typeParams} from "./types"

export { MsgVoteWeighted, MsgSubmitProposal, MsgDeposit, MsgUpdateParams, MsgVote };

type sendMsgVoteWeightedParams = {
  value: MsgVoteWeighted,
  fee?: StdFee,
  memo?: string
};

type sendMsgSubmitProposalParams = {
  value: MsgSubmitProposal,
  fee?: StdFee,
  memo?: string
};

type sendMsgDepositParams = {
  value: MsgDeposit,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateParamsParams = {
  value: MsgUpdateParams,
  fee?: StdFee,
  memo?: string
};

type sendMsgVoteParams = {
  value: MsgVote,
  fee?: StdFee,
  memo?: string
};


type msgVoteWeightedParams = {
  value: MsgVoteWeighted,
};

type msgSubmitProposalParams = {
  value: MsgSubmitProposal,
};

type msgDepositParams = {
  value: MsgDeposit,
};

type msgUpdateParamsParams = {
  value: MsgUpdateParams,
};

type msgVoteParams = {
  value: MsgVote,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgVoteWeighted({ value, fee, memo }: sendMsgVoteWeightedParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgVoteWeighted: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgVoteWeighted({ value: MsgVoteWeighted.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgVoteWeighted: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgSubmitProposal({ value, fee, memo }: sendMsgSubmitProposalParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSubmitProposal: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSubmitProposal({ value: MsgSubmitProposal.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSubmitProposal: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeposit({ value, fee, memo }: sendMsgDepositParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeposit: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeposit({ value: MsgDeposit.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeposit: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateParams({ value, fee, memo }: sendMsgUpdateParamsParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateParams: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateParams({ value: MsgUpdateParams.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateParams: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgVote({ value, fee, memo }: sendMsgVoteParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgVote: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgVote({ value: MsgVote.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgVote: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgVoteWeighted({ value }: msgVoteWeightedParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.gov.v1.MsgVoteWeighted", value: MsgVoteWeighted.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgVoteWeighted: Could not create message: ' + e.message)
			}
		},
		
		msgSubmitProposal({ value }: msgSubmitProposalParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.gov.v1.MsgSubmitProposal", value: MsgSubmitProposal.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSubmitProposal: Could not create message: ' + e.message)
			}
		},
		
		msgDeposit({ value }: msgDepositParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.gov.v1.MsgDeposit", value: MsgDeposit.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeposit: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateParams({ value }: msgUpdateParamsParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.gov.v1.MsgUpdateParams", value: MsgUpdateParams.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateParams: Could not create message: ' + e.message)
			}
		},
		
		msgVote({ value }: msgVoteParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.gov.v1.MsgVote", value: MsgVote.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgVote: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						WeightedVoteOption: getStructure(typeWeightedVoteOption.fromPartial({})),
						Deposit: getStructure(typeDeposit.fromPartial({})),
						Proposal: getStructure(typeProposal.fromPartial({})),
						TallyResult: getStructure(typeTallyResult.fromPartial({})),
						Vote: getStructure(typeVote.fromPartial({})),
						DepositParams: getStructure(typeDepositParams.fromPartial({})),
						VotingParams: getStructure(typeVotingParams.fromPartial({})),
						TallyParams: getStructure(typeTallyParams.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			CosmosGovV1: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;