import { useState } from 'react'
import {CreateGame} from "./axios.ts"
import './App.css'



function App() {
	
	const [code, setCode] = useState("");
	const [ruleset, setRuleset] = useState("Random");
	
	async function create_game() {
		let codes = await CreateGame(ruleset);
		console.log(codes);
		localStorage.setItem("guest_code", codes.guestCode);
		window.location.href = `/play?code=${codes.hostCode}`;
	}
	async function join_game() {
		window.location.href = `/play?code=${code}`;
	}

	return (
	<b>
		<label>
			Select a Ruleset: <br />
			<select name="ruleset" id="ruleset" onChange={(e) => setRuleset(e.target.value)}>
				<option value="Random">Random</option>
				<option value="Default">Default</option>
				<option value="Open World">Open World</option>
				<option value="Oops! All Knights!">Oops! All Knights!</option>
				<option value="PREPARE THYSELF">PREPARE THYSELF</option>
				<option value="Have a plan to kill everyone you meet">Have a plan to kill everyone you meet</option>
				<option value="Atomic Chess">Atomic Chess</option>
				<option value="Knook">Knook</option>
			</select>
		</label>
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
