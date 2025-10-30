import axios from 'axios';
import { ResponseGame } from './models/ResponseGame';
import { Status } from './models/Status';
import NewGameResponse from './models/NewGameResponse';
import { ResponseRulesets } from './models/ResponseRulesets';

const backend_url: string = import.meta.env.VITE_REACT_APP_BACKEND_URL.endsWith("/") ? (import.meta.env.VITE_REACT_APP_BACKEND_URL as string).slice(0, -1) : import.meta.env.VITE_REACT_APP_BACKEND_URL

export async function CreateGame(ruleName: string): Promise<NewGameResponse> {
	const url = `${backend_url}/game/new`;
	const response = await axios.post(url, `{"ruleName": "${ruleName}"}`);
	return response.data as NewGameResponse;
}

export async function Fetch(code: string): Promise<[ResponseGame, string]> {
	const url = `${backend_url}/game/fetch`;
	const response = await axios.post(url, `{"code": "${code}"}`, {validateStatus: (_) => {return true}});

	if (response.status != 200) {
		return [{} as ResponseGame, "Invalid Code"]
	}

	return [response.data as ResponseGame, ""];
}

export async function Move(code: string, start_pos: number, end_pos: number): Promise<Status> {
	const url = `${backend_url}/game/move`;
	const response = await axios.post(url, `{"code": "${code}", "start_pos": ${start_pos}, "end_pos": ${end_pos}}`);
	return response.data as Status;
}

export async function FetchRulesets(): Promise<ResponseRulesets> {
	const url = `${backend_url}/info/rulesets`;
	const response = await axios.get(url);
	return response.data as ResponseRulesets;
}