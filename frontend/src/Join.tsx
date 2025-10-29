import { useEffect, useState } from 'react'
import {CreateGame, FetchRulesets} from "./axios.ts"
import './App.css'



function App() {
	
	const [code, setCode] = useState("");
	const [ruleset, setRuleset] = useState("Random");
	const [rulesets, setRulesets] = useState<JSX.Element[]>([])
	
	async function create_game() {
		let codes = await CreateGame(ruleset);
		console.log(codes);
		localStorage.setItem("guest_code", codes.guestCode);
		window.location.href = `/play?code=${codes.hostCode}`;
	}
	async function join_game() {
		window.location.href = `/play?code=${code}`;
	}

	async function get_rulesets(): Promise<JSX.Element[]> {
		let rulesets = await FetchRulesets()

		rulesets.rulesets = rulesets.rulesets.sort()

		rulesets.rulesets = ["Random"].concat(rulesets.rulesets)

		let html: JSX.Element[] = []

		rulesets.rulesets.forEach(element => {
			if (element == "Default") {
				html = [html[0], <option value={element}>{element}</option>].concat(html.slice(1))
				return
			}
			html.push(
				<option value={element}>{element}</option>
			)
		});

		return html
	}

	useEffect(() => {
		async function init() {
			setRulesets(await get_rulesets())
		}
		init()
	}, [])

	return (
	<b>
		<label>
			Select a Ruleset: <br />
			<select name="ruleset" id="ruleset" onChange={(e) => setRuleset(e.target.value)}>
				{rulesets}
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
