import axios from 'axios';

export async function CreateGame() {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/new`;
	const response = await axios.get(url);
	return response.data;
}

export async function JoinGame(code: string) {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/exists/${code}`;
	const response = await axios.get(url);
	return response.data;
}

export async function Fetch(code: string) {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/fetch/${code}`;
	const response = await axios.get(url);
	return response.data;
}

export async function Move(code: string, start_pos: number, end_pos: number) {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/move/${code}`;
	const response = await axios.post(url, `${start_pos}\n${end_pos}`);
	return response.data;
}

export async function Leave(code: string) {
	const url = `${import.meta.env.VITE_REACT_APP_BACKEND_URL}/game/leave/${code}`;
	const response = await axios.get(url);
	return response.data;
}
