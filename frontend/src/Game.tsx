//import font from '/public/3270NerdFont-Regular.ttf'
import { useState, useEffect } from 'react'
import { Fetch, Move } from './axios.ts'
import './Game.css'
import { useSearchParams } from 'react-router-dom'
import Team from './models/Team.ts';


function Game() {
	const [chessBoard, setChessBoard] = useState<JSX.Element[]>([]);
	const [boardData, setBoardData] = useState<number[]>([]);
	const [boardDimensions, setBoardDimensions] = useState<[number, number]>([8,8]);
	const [turn, setTurn] = useState<Team>(Team.White);
	const [rule, setRule] = useState<string>("");
	const [team, setTeam] = useState<Team>(Team.NoTeam);
	const [start_pos, setStart_pos] = useState<number>(-1);
	const [end_pos, setEnd_pos] = useState<number>(-1);
	const [winner, setWinner] = useState<Team>(0);

	let [query] = useSearchParams();
	let pieces: string[] = ["", "", "", "", "", ""];
	let pieceNames: string[] = ["pawn", "rook", "knight", "bishop", "queen", "king"]
	
	function manageClick(number: number) {
		if(team == Team.NoTeam) {
			return
		}

		if(turn != team) {
			return
		}

		if(start_pos == -1) {
			setStart_pos(number);
		} else if (end_pos == -1) {
			setEnd_pos(number);
		} else {
			setStart_pos(-1);
			setEnd_pos(-1);
		}
	}

	function RenderSquare(i: number) {
		let color: string = "whitetext";
		let active: boolean = false
		let piece: number = boardData[i];

		// ignore the "has moved" bit
		if (piece >= 32) {
			piece -= 32
		}
		
		if(piece >= 16) {
			piece -= 16
			active = true
		}

		if(piece >= 8) {
			piece -= 8;
			color = "blacktext";
		}

		if (piece > 6) {
			console.log(piece)
		}

		let column = Math.floor(i/boardDimensions[0])
		let classname: string = (i+column) % 2 == 0 ? "square light" : "square dark";

		return <div className={start_pos == i || end_pos == i ? "square red" : classname} key={"board element "+i} onClick={() => {if (winner == Team.NoTeam) {manageClick(i)} }}>
			{active ? <label className={`${color} ${pieceNames[piece]}`}>{pieces[piece]}</label> : <></>}
		</div>
	}

	useEffect(() => {
		if(winner != Team.NoTeam) {
			alert(`${winner == 2 ? "White" : "Black"} Won`)
			return
		}

		let squares: JSX.Element[] = []
		
		for(let i: number = 0; i < boardDimensions[0]*boardDimensions[1]; i++) {
			if(i%boardDimensions[0] == 0 && i != 0) {
				squares.push(<br key={"break "+i}/>);
			}
			squares.push(RenderSquare(i));
		}

		if(team == Team.Black) {
			squares.reverse();
		}

		setChessBoard(squares);
	}, [boardData, start_pos, end_pos]);
	
	useEffect(() => {
		async function run() {
			if(start_pos == -1 || end_pos == -1) {
				return;
			}

			let code = query.get("code")
			if(code == null) {
				code = "";
			}
			if ((await Move(code, start_pos, end_pos)).status == "success") {

				//do the move client side so its more responsive
				let newBoard: number[] = boardData;
				newBoard[end_pos] = newBoard[start_pos];
				newBoard[start_pos] = 0;
				setBoardData(newBoard);

				setStart_pos(-1);
				setEnd_pos(-1);
				return;
			}
		}
		if(winner == Team.NoTeam) {
			run();
		}
	}, [start_pos, end_pos])
	
	useEffect(() => {
		let interval = setInterval(async() => {
			if(winner != 0) {
				return
			}

			let code = query.get("code")
			if(code == null) {
				code = "";
			}
			let [fetch, error] = await Fetch(code)
			if (error != "") {
				alert("Invalid Game Code")
				window.location.href = "/"
			}

			setBoardData(fetch.game.board.pieces);
			setBoardDimensions([fetch.game.board.width, fetch.game.board.height]);
			setTeam(fetch.team)
			setWinner(fetch.game.winner);
			setTurn(fetch.game.turn);
			setRule(fetch.game.ruleset.name);
		}, 1000);
		

		return () => clearInterval(interval);
	}, [boardData]);
	
	async function leaveGame() {
		localStorage.removeItem("guest_code")
		window.location.href = "/";
	}
	
	return (
	<b>
		{chessBoard}
		<div className="alignBottom">
			<button className="bottomElement red" onClick={() => {leaveGame()}}>Leave</button>
			<label className="bottomElement grey">Join Code: {localStorage.getItem("guest_code")}</label>
			<label className="bottomElement grey">Rule: {rule}</label>
		</div>
	</b>
	)
}

export default Game;
