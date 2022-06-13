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
    makeDelete,
    createDeleteAsync,
    resetMessage,

} from '../banking/BankingSlice';
import styles from '../banking/Banking.module.css';
import {Box, TextField} from "@mui/material";

export function Account() {
    const banking = useAppSelector(selectBanking);
    const bankingUser = useAppSelector(selectBankingUser)
    const dispatch = useAppDispatch();
    const [username, setUsername] = useState('');
    const [amount, setStateAmount] = useState('0');
    const [password, setPassword] = useState('');
    const [display, setDisplay] = useState(true);
    const [displayUserCreation, setUserCreation] = useState(false);
    const [displayLoginCreation, setLoginCreation] = useState(false);
    const [displayDeleteCreation, setDeleteCreation] = useState(false);
    const amountValue = Number(amount) || 0;

    let createInfoElem =
        <div className={styles.row}>
            <button
                className={styles.button}
                onClick={() => createDelete(dispatch, username)}
            >
                Delete Account
            </button>

        </div>;

    return (
        <div>
            <div className={styles.row} id="welcomeDiv">
                {createInfoElem}
            </div>
        </div>
    );
}

function createDelete(dispatch: any, account: string) {
    dispatch(makeDelete({account}))
    dispatch(createDeleteAsync({account}))
}

function openDeleteCreation(setDisplay: any, setDeleteCreation: any) {
    setDisplay(false)
    setDeleteCreation(true)
}

function closeDeleteCreation(setDisplay: any, setDeleteCreation: any) {
    setDisplay(true)
    setDeleteCreation(false)
}