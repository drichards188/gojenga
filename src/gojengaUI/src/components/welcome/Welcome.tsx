import React, {useState} from 'react';

import {useAppSelector, useAppDispatch} from '../../app/hooks';
import {
    decrement,
    increment,
    setUser,
    createUser,
    setAmount,
    incrementByAmount,
    incrementAsync,
    createUserAsync,
    incrementIfOdd,
    selectBanking,
    selectBankingUser,
    makeTransaction,
    createTransactionAsync,
    makeDeposit,
    makeLogin,
    createDepositAsync,
    createDeleteAsync,
    makeDelete,
    makeInfo,
    createInfoAsync, createLoginAsync

} from '../banking/BankingSlice';
import styles from '../banking/Banking.module.css';
import {Counter} from "../counter/Counter";

// @ts-ignore
export function Welcome() {
    const banking = useAppSelector(selectBanking);
    const bankingUser = useAppSelector(selectBankingUser)
    const dispatch = useAppDispatch();
    const [incrementAmount, setIncrementAmount] = useState('200');
    const [username, setUsername] = useState('');
    const [destination, setDestination] = useState('');
    const [amount, setStateAmount] = useState('0');
    const [password, setPassword] = useState('');
    const [display, setDisplay] = useState(true);
    const [displayUserCreation, setUserCreation] = useState(false);
    const [displayTransactionCreation, setTransactionCreation] = useState(false);
    const [displayDepositCreation, setDepositCreation] = useState(false);
    const [displayInfoCreation, setInfoCreation] = useState(false);
    const [displayDeleteCreation, setDeleteCreation] = useState(false);
    const [displayLoginCreation, setLoginCreation] = useState(false);
    const incrementValue = Number(incrementAmount) || 0;
    const amountValue = Number(amount) || 0;

    let output;
    if (display) {
        output =
            <div className={styles.row}>
                <div>

                    <button
                        className={styles.button}
                        onClick={() => openAccountCreation(setDisplay, setUserCreation)}
                    >
                        Create Account
                    </button>
                    <button
                        className={styles.button}
                        onClick={() => openLoginCreation(setDisplay, setLoginCreation)}
                    >
                        Login
                    </button>
                </div>
            </div>
    }

    let createUserElem;
    if (displayUserCreation) {

        createUserElem =
            <div className={styles.row}>
                <div>
                    <input
                        className={styles.textbox}
                        aria-label="Set User"
                        placeholder={"Username"}
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <input
                        className={styles.textbox}
                        aria-label="Set Amount"
                        value={amountValue}
                        onChange={(e) => setStateAmount(e.target.value)}
                    />
                    {/*<span className={styles.value}>{banking}</span>*/}
                    {/*<span className={styles.value} onClick={() => dispatch(setUser("Tucker"))}>{blockchainUser}</span>*/}
                </div>
                <button
                    className={styles.button}
                    onClick={() => createMyUser(dispatch, username, amount)}
                >
                    Create User
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeAccountCreation(setDisplay, setUserCreation)}
                >
                    Back
                </button>
            </div>;
        // setTransactionCreation(false)
    }

    let createLoginElem;
    if (displayLoginCreation) {

        createLoginElem =
            <div className={styles.row}>
                <div>
                    <input
                        className={styles.textbox}
                        aria-label="Set User"
                        placeholder={"Username"}
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <input
                        className={styles.textbox}
                        aria-label="Set Password"
                        placeholder={"Password"}
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    {/*<span className={styles.value}>{banking}</span>*/}
                    {/*<span className={styles.value} onClick={() => dispatch(setUser("Tucker"))}>{blockchainUser}</span>*/}
                </div>
                <button
                    className={styles.button}
                    onClick={() => createLogin(dispatch, username, password)}
                >
                    Login
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeLoginCreation(setDisplay, setLoginCreation)}
                >
                    Back
                </button>
            </div>;
        // setTransactionCreation(false)
    }

    // @ts-ignore
    return (
        <div>
            <div className={styles.row} id="welcomeDiv">
                {createUserElem}
                {createLoginElem}
            </div>
            {output}
        </div>
    );
}

function createMyUser(dispatch: any, username: any, amount:any) {
    dispatch(createUser({username, amount}))
    dispatch(createUserAsync({username, amount}))
}


function createLogin(dispatch: any, account:any, password: any) {
    dispatch(makeLogin({account, password}))
    dispatch(createLoginAsync({account, password}))
}

function openAccountCreation(setDisplay: any, setUserCreation: any) {
    setDisplay(false)
    setUserCreation(true)
}

function openLoginCreation(setDisplay: any, setLoginCreation: any) {
    setDisplay(false)
    setLoginCreation(true)
}

function closeAccountCreation(setDisplay: any, setUserCreation: any) {
    setDisplay(true)
    setUserCreation(false)
}

function closeLoginCreation(setDisplay: any, setLoginCreation: any) {
    setDisplay(true)
    setLoginCreation(false)
}