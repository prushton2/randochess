import axios from 'axios';
import { ResponseGame } from './models/ResponseGame';
import { Status } from './models/Status';

export async function CreateGame() {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/new`;
	const response = await axios.get(url);
	return response.data;
}

export async function JoinGame(code: string) {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/exists`;
	const response = await axios.post(url, `{"code": "${code}"}`);
	return response.data;
}

export async function Fetch(code: string): Promise<ResponseGame> {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/fetch`;
	const response = await axios.post(url, `{"code": "${code}"}`);
	return response.data as ResponseGame;
}

export async function Move(code: string, start_pos: number, end_pos: number): Promise<Status> {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/move`;
	const response = await axios.post(url, `{"code": "${code}", "start_pos": ${start_pos}, "end_pos": ${end_pos}}`);
	return response.data as Status;
}

export async function Leave(code: string) {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/leave`;
	const response = await axios.post(url, `{"code": "${code}"}`);
	return response.data;
}
