import React, {useState} from 'react';

import {useAppSelector, useAppDispatch} from '../../app/hooks';
import {
    createUser,
    createUserAsync,
    selectBanking,
    selectBankingUser,
    makeLogin,
    createLoginAsync,
    makeInfo,
    createInfoAsync,
    makeDeposit,
    createDepositAsync,
    makeTransaction,
    createTransactionAsync,

} from '../banking/BankingSlice';
import styles from '../banking/Banking.module.css';
import {Box, TextField} from "@mui/material";

export function Transaction() {
    const banking = useAppSelector(selectBanking);
    const bankingUser = useAppSelector(selectBankingUser)
    const dispatch = useAppDispatch();
    const [username, setUsername] = useState('');
    const [amount, setStateAmount] = useState('0');
    const [destination, setDestination] = useState('');
    const [display, setDisplay] = useState(true);
    const [displayUserCreation, setUserCreation] = useState(false);
    const [displayLoginCreation, setLoginCreation] = useState(false);
    const amountValue = Number(amount) || 0;

    let createTransactionElem =
        <div className={styles.row}>
            <div>
                <TextField
                    id="destination"
                    label="Destination"
                    variant="standard"
                    className={styles.textbox}
                    aria-label="Set User"
                    placeholder={"Destination Username"}
                    value={destination}
                    onChange={(e) => setDestination(e.target.value)}
                />
                <TextField
                    id="payment-amount"
                    label="Payment Amount"
                    variant="standard"
                    type="number"
                    inputMode="numeric"
                    className={styles.textbox}
                    aria-label="Pay Amount"
                    value={amountValue}
                    onChange={(e) => setStateAmount(e.target.value)}
                />
            </div>
            <button
                className={styles.button}

                onClick={() => createTransaction(dispatch, bankingUser, destination, amount)}
            >
                Pay
            </button>
        </div>;

    return (
        <div>
            <div className={styles.row} id="welcomeDiv">
                {createTransactionElem}
            </div>
        </div>
    );
}

function createTransaction(dispatch: any, account: string, destination: string, amount: string) {
    dispatch(makeTransaction({destination, amount}))
    dispatch(createTransactionAsync({account, destination, amount}))
}


