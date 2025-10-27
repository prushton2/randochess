import { useState } from 'react'
import {CreateGame, JoinGame} from "./axios.ts"
import './App.css'



function App() {
	
	const [code, setCode] = useState("");
	
	async function create_game() {
		let codes = await CreateGame();
		console.log(codes);
		localStorage.setItem("guest_code", codes.guest_code);
		window.location.href = `/play?code=${codes.host_code}`;
	}
	async function join_game() {
		let result: string = await JoinGame(code);
		if(result == "1") {
			window.location.href = `/play?code=${code}`;
		} else {
			alert("Invalid Code");
		}
	}

	return (
	<b>
		<button onClick={create_game}>
		      Create Game
		</button>
		<br/><br/>
		or
		<br/><br/>
		<input type="number" onChange={(e) => {setCode(e.target.value)}}/>

		<button onClick={join_game}>
			Join
		</button>
	</b>
	)
}

export default App
