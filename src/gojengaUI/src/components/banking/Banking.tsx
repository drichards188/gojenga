import React, {useState} from 'react';

import {Welcome} from "../welcome/Welcome";
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
    selectLoggedIn,
    makeTransaction,
    createTransactionAsync,
    makeDeposit,
    makeLogin,
    createDepositAsync,
    createDeleteAsync,
    makeDelete,
    makeInfo,
    createInfoAsync, createLoginAsync

} from './BankingSlice';
import styles from './Banking.module.css';
import {Counter} from "../counter/Counter";

// @ts-ignore
export function Banking() {
    const banking = useAppSelector(selectBanking);
    const bankingUser = useAppSelector(selectBankingUser)
    const isLoggedIn = useAppSelector(selectLoggedIn)
    const dispatch = useAppDispatch();
    const [incrementAmount, setIncrementAmount] = useState('200');
    const [username, setUsername] = useState('');
    const [destination, setDestination] = useState('');
    const [amount, setStateAmount] = useState('0');
    const [password, setPassword] = useState('');
    const [display, setDisplay] = useState(false);
    const [displayUserCreation, setUserCreation] = useState(false);
    const [displayTransactionCreation, setTransactionCreation] = useState(false);
    const [displayDepositCreation, setDepositCreation] = useState(false);
    const [displayInfoCreation, setInfoCreation] = useState(false);
    const [displayDeleteCreation, setDeleteCreation] = useState(false);
    const [displayLoginCreation, setLoginCreation] = useState(false);
    const incrementValue = Number(incrementAmount) || 0;
    const amountValue = Number(amount) || 0;

    let output;
    if (isLoggedIn) {
        output =
            <div className={styles.row}>
                <div>
                    <button
                        className={styles.button}
                        onClick={() => openDepositCreation(setDisplay, setDepositCreation)}
                    >
                        Deposit
                    </button>
                    <button
                        className={styles.button}
                        onClick={() => openTransactionCreation(setDisplay, setTransactionCreation)}
                    >
                        Pay
                    </button>
                    <button
                        className={styles.button}
                        onClick={() => openInfoCreation(setDisplay, setInfoCreation)}
                    >
                        Balance
                    </button>
                    <button
                        className={styles.button}
                        onClick={() => openDeleteCreation(setDisplay, setDeleteCreation)}
                    >
                        Delete Account
                    </button>
                </div>
            </div>
    }

    let welcomeElem;
    if (!isLoggedIn) {
        welcomeElem =
            <Welcome/>
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

    let createTransactionElem;
    if (displayTransactionCreation) {

        createTransactionElem =
            <div className={styles.row}>
                <div>
                    <input
                        className={styles.textbox}
                        aria-label="Set User"
                        placeholder={"Destination Username"}
                        value={destination}
                        onChange={(e) => setDestination(e.target.value)}
                    />
                    <input
                        className={styles.textbox}
                        aria-label="Pay Amount"
                        value={amountValue}
                        onChange={(e) => setStateAmount(e.target.value)}
                    />
                    {/*<span className={styles.value}>{banking}</span>*/}
                    {/*<span className={styles.value} onClick={() => dispatch(setUser("Tucker"))}>{blockchainUser}</span>*/}
                </div>
                <button
                    className={styles.button}
                    onClick={() => createTransaction(dispatch, username, destination, amount)}
                >
                    Pay
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeTransactionCreation(setDisplay, setTransactionCreation)}
                >
                    Back
                </button>
            </div>;
        // setUserCreation(false)
    }

    let createDepositElem;
    if (displayDepositCreation) {

        createDepositElem =
            <div className={styles.row}>
                <div>
                    <input
                        className={styles.textbox}
                        aria-label="Set User"
                        placeholder={"Username"}
                        value={bankingUser}
                        // onChange={(e) => setUsername(e.target.value)}
                    />
                    <input
                        className={styles.textbox}
                        aria-label="Set User"
                        placeholder={"Username"}
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <input
                        className={styles.textbox}
                        aria-label="Deposit Amount"
                        value={amountValue}
                        onChange={(e) => setStateAmount(e.target.value)}
                    />
                    {/*<span className={styles.value}>{banking}</span>*/}
                    {/*<span className={styles.value} onClick={() => dispatch(setUser("Tucker"))}>{blockchainUser}</span>*/}
                </div>
                <button
                    className={styles.button}
                    onClick={() => createDeposit(dispatch, username, amount)}
                >
                    Deposit
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeDepositCreation(setDisplay, setDepositCreation)}
                >
                    Back
                </button>
            </div>;
        // setUserCreation(false)
    }

    let createInfoElem;
    if (displayInfoCreation) {

        createInfoElem =
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
                        aria-label="Deposit Amount"
                        value={banking}
                    />
                    {/*<span className={styles.value}>{banking}</span>*/}
                    {/*<span className={styles.value} onClick={() => dispatch(setUser("Tucker"))}>{blockchainUser}</span>*/}
                </div>
                <button
                    className={styles.button}
                    onClick={() => createInfo(dispatch, username)}
                >
                    See Balance
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeInfoCreation(setDisplay, setInfoCreation)}
                >
                    Back
                </button>
            </div>;
        // setUserCreation(false)
    }

    let createDeleteElem;
    if (displayDeleteCreation) {

        createDeleteElem =
            <div className={styles.row}>
                <div>
                    <input
                        className={styles.textbox}
                        aria-label="Set User"
                        placeholder={"Username"}
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    {/*<span className={styles.value}>{banking}</span>*/}
                    {/*<span className={styles.value} onClick={() => dispatch(setUser("Tucker"))}>{blockchainUser}</span>*/}
                </div>
                <button
                    className={styles.button}
                    onClick={() => createDelete(dispatch, username)}
                >
                    Delete Account
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeDepositCreation(setDisplay, setDeleteCreation)}
                >
                    Back
                </button>
            </div>;
        // setUserCreation(false)
    }

    // @ts-ignore
    return (
        <div>
            <div className={styles.row}>
                {createTransactionElem}
                {createDepositElem}
                {createInfoElem}
                {createDeleteElem}
            </div>
            {welcomeElem}
            {output}
        </div>
    );
}

function createMyUser(dispatch: any, username: any, amount: any) {
    dispatch(createUser({username, amount}))
    dispatch(createUserAsync({username, amount}))
}

function createTransaction(dispatch: any, account: any, destination: any, amount: any) {
    dispatch(makeTransaction({destination, amount}))
    dispatch(createTransactionAsync({account, destination, amount}))
}

function createLogin(dispatch: any, account: any, password: any) {
    dispatch(makeLogin({account, password}))
    dispatch(createLoginAsync({account, password}))
}

function createDeposit(dispatch: any, account: any, amount: any) {
    dispatch(makeDeposit({account, amount}))
    dispatch(createDepositAsync({account, amount}))
}

function createInfo(dispatch: any, account: any) {
    dispatch(makeInfo({account}))
    dispatch(createInfoAsync({account}))
}

function createDelete(dispatch: any, account: any) {
    dispatch(makeDelete({account}))
    dispatch(createDeleteAsync({account}))
}

function openAccountCreation(setDisplay: any, setUserCreation: any) {
    setDisplay(false)
    setUserCreation(true)
}

function openLoginCreation(setDisplay: any, setLoginCreation: any) {
    setDisplay(false)
    setLoginCreation(true)
}

function openTransactionCreation(setDisplay: any, setTransactionCreation: any) {
    setDisplay(false)
    setTransactionCreation(true)
}

function openDepositCreation(setDisplay: any, setDepositCreation: any) {
    setDisplay(false)
    setDepositCreation(true)
}

function openInfoCreation(setDisplay: any, setInfoCreation: any) {
    setDisplay(false)
    setInfoCreation(true)
}

function openDeleteCreation(setDisplay: any, setDeleteCreation: any) {
    setDisplay(false)
    setDeleteCreation(true)
}

function closeAccountCreation(setDisplay: any, setUserCreation: any) {
    setDisplay(true)
    setUserCreation(false)
}

function closeTransactionCreation(setDisplay: any, setTransactionCreation: any) {
    setDisplay(true)
    setTransactionCreation(false)
}

function closeDepositCreation(setDisplay: any, setDepositCreation: any) {
    setDisplay(true)
    setDepositCreation(false)
}

function closeInfoCreation(setDisplay: any, setInfoCreation: any) {
    setDisplay(true)
    setInfoCreation(false)
}

function closeDeleteCreation(setDisplay: any, setDeleteCreation: any) {
    setDisplay(true)
    setDeleteCreation(false)
}

function closeLoginCreation(setDisplay: any, setLoginCreation: any) {
    setDisplay(true)
    setLoginCreation(false)
}