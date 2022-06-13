import React, {useState} from 'react';

import {Welcome} from "../welcome/Welcome";
import {useAppSelector, useAppDispatch} from '../../app/hooks';
import {
    createUser,
    createUserAsync,
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
    resetMessage,
    makeInfo,
    createInfoAsync, createLoginAsync, selectToken, selectMessage, selectBalance, selectAmount

} from './BankingSlice';
import styles from './Banking.module.css';

export function Banking() {
    const dispatch = useAppDispatch();
    const banking = useAppSelector(selectBanking);
    const bankingUser = useAppSelector(selectBankingUser);
    const token = useAppSelector(selectToken);
    const serverMessage = useAppSelector(selectMessage);
    const balance = useAppSelector(selectBalance);
    const serverAmount = useAppSelector(selectAmount);
    const isLoggedIn = useAppSelector(selectLoggedIn);


    const [username, setUsername] = useState('');
    const [destination, setDestination] = useState('');
    const [amount, setStateAmount] = useState('0');
    const [display, setDisplay] = useState(true);
    const [displayTransactionCreation, setTransactionCreation] = useState(false);
    const [displayDepositCreation, setDepositCreation] = useState(false);
    const [displayInfoCreation, setInfoCreation] = useState(false);
    const [displayDeleteCreation, setDeleteCreation] = useState(false);
    const amountValue = Number(amount) || 0;

    let toolbar;
    if (isLoggedIn && display) {
        toolbar =

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
                        onClick={() => openInfoCreation(setDisplay, setInfoCreation, dispatch, username)}
                    >
                        Account
                    </button>
                </div>
            </div>
    }

    let welcomeElem;
    if (!isLoggedIn) {
        welcomeElem =
            <Welcome/>
    }

    let createTransactionElem;
    if (displayTransactionCreation) {

        createTransactionElem =
            <div className={styles.row}>
                <text
                    className={styles.textbox}
                    aria-label="Set User"
                >
                    {serverMessage}
                </text>
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
                </div>
                <button
                    className={styles.button}
                    onClick={() => createTransaction(dispatch, username, destination, amount)}
                >
                    Pay
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeTransactionCreation(setDisplay, setTransactionCreation, dispatch)}
                >
                    Back
                </button>
            </div>;
    }

    let createDepositElem;
    if (displayDepositCreation) {

        createDepositElem =
            <div className={styles.row}>
                <div>
                    <text
                        className={styles.textbox}
                        aria-label="Set User"
                    >
                        {serverMessage}
                    </text>
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
                </div>
                <button
                    className={styles.button}
                    onClick={() => createDeposit(dispatch, username, amount)}
                >
                    Deposit
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeDepositCreation(setDisplay, setDepositCreation, dispatch)}
                >
                    Back
                </button>
            </div>;
    }

    let createInfoElem;
    if (displayInfoCreation) {

        createInfoElem =
            <div className={styles.row}>
                <div>
                    <text
                        className={styles.textbox}
                        aria-label="Set User"
                    >
                        {bankingUser}
                    </text>
                    <text
                        className={styles.textbox}
                        aria-label="Set User"
                    >
                        {serverAmount}
                    </text>
                </div>
                <button
                    className={styles.button}
                    onClick={() => openDeleteCreation(setDisplay, setDeleteCreation)}
                >
                    Delete Account
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeInfoCreation(setDisplay, setInfoCreation, dispatch)}
                >
                    Back
                </button>
            </div>;
    }

    let createDeleteElem;
    if (displayDeleteCreation) {

        createDeleteElem =
            <div className={styles.row}>
                <text
                    className={styles.textbox}
                    aria-label="Set User"
                >
                    {serverMessage}
                </text>
                <div>
                    <input
                        className={styles.textbox}
                        aria-label="Set User"
                        placeholder={"Username"}
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </div>
                <button
                    className={styles.button}
                    onClick={() => createDelete(dispatch, username)}
                >
                    Delete Account
                </button>
                <button
                    className={styles.button}
                    onClick={() => closeDepositCreation(setDisplay, setDeleteCreation, dispatch)}
                >
                    Back
                </button>
            </div>;
    }

    let infoDiv;
    if (isLoggedIn && display) {
        infoDiv =
            <div>
                <p
                    className={styles.textbox}
                    aria-label="Set User"
                >
                    {"Hi "  + bankingUser + "!"}
                </p>
                <p
                    className={styles.textbox}
                    aria-label="Set User"
                >
                    {"You have $" + balance}
                </p>
            </div>
    }

    return (
        <div>
            {infoDiv}
            <div className={styles.row}>
                {createTransactionElem}
                {createDepositElem}
                {createInfoElem}
                {createDeleteElem}
            </div>
            {welcomeElem}
            {toolbar}
        </div>
    );
}

function createTransaction(dispatch: any, account: string, destination: string, amount: string) {
    dispatch(makeTransaction({destination, amount}))
    dispatch(createTransactionAsync({account, destination, amount}))
}

function createDeposit(dispatch: any, account: string, amount: string) {
    dispatch(makeDeposit({account, amount}))
    dispatch(createDepositAsync({account, amount}))
}

function createInfo(dispatch: any, account: string) {
    dispatch(makeInfo({account}))
    dispatch(createInfoAsync({account}))
}

function createDelete(dispatch: any, account: string) {
    dispatch(makeDelete({account}))
    dispatch(createDeleteAsync({account}))
}

function openTransactionCreation(setDisplay: any, setTransactionCreation: any) {
    setDisplay(false)
    setTransactionCreation(true)
}

function openDepositCreation(setDisplay: any, setDepositCreation: any) {
    setDisplay(false)
    setDepositCreation(true)
}

function openInfoCreation(setDisplay: any, setInfoCreation: any, dispatch:any, username: string) {
    setDisplay(false)
    setInfoCreation(true)
    createInfo(dispatch, username)
}

function openDeleteCreation(setDisplay: any, setDeleteCreation: any) {
    setDisplay(false)
    setDeleteCreation(true)
}

function closeTransactionCreation(setDisplay: any, setTransactionCreation: any, dispatch: any) {
    setDisplay(true)
    setTransactionCreation(false)
    dispatch(resetMessage())
}

function closeDepositCreation(setDisplay: any, setDepositCreation: any, dispatch: any) {
    setDisplay(true)
    setDepositCreation(false)
    dispatch(resetMessage())
}

function closeInfoCreation(setDisplay: any, setInfoCreation: any, dispatch: any) {
    setDisplay(true)
    setInfoCreation(false)
    dispatch(resetMessage())
}

function closeDeleteCreation(setDisplay: any, setDeleteCreation: any) {
    setDisplay(true)
    setDeleteCreation(false)
}