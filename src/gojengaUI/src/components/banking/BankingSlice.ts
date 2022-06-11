import {createAsyncThunk, createReducer, createSlice, PayloadAction, unwrapResult} from '@reduxjs/toolkit';
import {RootState, AppThunk} from '../../app/store';
import {crtDelete, crtDeposit, crtInfo, crtLogin, crtTransaction, crtUser, fetchCount} from './BankingAPI';

export interface BankingState {
    amount: number;
    user: string;
    password: string;
    destination: string;
    message: string;
    loggedIn: boolean;
    status: 'idle' | 'loading' | 'failed';
}

const initialState: BankingState = {
    amount: 0,
    user: 'david',
    password: '12345',
    destination: 'allie',
    message: 'msg',
    loggedIn: false,
    status: 'idle'
};

// The function below is called a thunk and allows us to perform async logic. It
// can be dispatched like a regular action: `dispatch(incrementAsync(10))`. This
// will call the thunk with the `dispatch` function as the first argument. Async
// code can then be executed and other actions can be dispatched. Thunks are
// typically used to make async requests.
export const incrementAsync = createAsyncThunk(
    'banking/incrementAsync',
    async (amount: number) => {
        const response = await fetchCount(amount);
        // The value we return becomes the `fulfilled` action payload

        return response.data;
    }
);

export const createUserAsync = createAsyncThunk(
    'banking/createUser',
    async (payload: any) => {
        const response = await crtUser(payload.username, payload.amount);
        // The value we return becomes the `fulfilled` action payload

        return response.data;
    }
);

export const createTransactionAsync = createAsyncThunk(
    'banking/createTransaction',
    async (payload: any) => {
        const response = await crtTransaction(payload.account, payload.destination, payload.amount);
        // The value we return becomes the `fulfilled` action payload

        return response.data;
    }
);

export const createLoginAsync = createAsyncThunk(
    'banking/createLogin',
    async (payload: any) => {
        const response = await crtLogin(payload.account, payload.password);
        // The value we return becomes the `fulfilled` action payload

        return response.data;
    }
);

export const createDepositAsync = createAsyncThunk(
    'banking/createDeposit',
    async (payload: any) => {
        const response = await crtDeposit(payload.account, payload.amount);
        // The value we return becomes the `fulfilled` action payload

        return response.data;
    }
);

export const createInfoAsync = createAsyncThunk(
    'banking/createInfo',
    async (payload: any) => {
        const response = await crtInfo(payload.account);
        // The value we return becomes the `fulfilled` action payload

        return response.data;
    }
);

export const createDeleteAsync = createAsyncThunk(
    'banking/createDelete',
    async (payload: any) => {
        const response = await crtDelete(payload.account);
        // The value we return becomes the `fulfilled` action payload

        return response.data;
    }
);

export const bankingSlice = createSlice({
    name: 'banking',
    initialState,
    // The `reducers` field lets us define reducers and generate associated actions
    reducers: {
        increment: (state) => {
            // Redux Toolkit allows us to write "mutating" logic in reducers. It
            // doesn't actually mutate the state because it uses the Immer library,
            // which detects changes to a "draft state" and produces a brand new
            // immutable state based off those changes
            state.amount += 1;
        },
        decrement: (state) => {
            state.amount -= 1;
        },
        // Use the PayloadAction type to declare the contents of `action.payload`
        incrementByAmount: (state, action: PayloadAction<number>) => {
            state.amount += action.payload;
        },
        setUser: (state, action: PayloadAction<string>) => {
            state.user = action.payload
            alert("setUser state is " + state.user)
        },
        setAmount: (state, action: PayloadAction<number>) => {
            state.amount = action.payload
        },
        resetState: (state) => {
            return initialState
        },
        createUser: (state, action: PayloadAction<any>) => {
            state.user = action.payload.username
            state.amount = action.payload.amount
        },
        makeLogin: (state, action: PayloadAction<any>) => {
            state.user = action.payload.username
            state.password = action.payload.password
        },
        makeTransaction: (state, action: PayloadAction<any>) => {
            state.user = action.payload.username
            state.amount = action.payload.amount
        },
        makeDeposit: (state, action: PayloadAction<any>) => {
            state.user = action.payload.username
            state.amount = action.payload.amount
        },
        makeInfo: (state, action: PayloadAction<any>) => {
            state.user = action.payload.username
        },
        makeDelete: (state, action: PayloadAction<any>) => {
            state.user = action.payload.username
            state.amount = action.payload.amount
        },
    },
    // The `extraReducers` field lets the slice handle actions defined elsewhere,
    // including actions generated by createAsyncThunk or in other slices.
    //todo state data should be changing correctly. should only need to call state data in ui to reflect response
    extraReducers: (builder) => {
        builder
            //createTransaction
            .addCase(createTransactionAsync.pending, (state) => {
                state.status = 'loading';
            })
            .addCase(createTransactionAsync.fulfilled, (state, action) => {
                state.status = 'idle';
                state.user = action.payload["error"]["name"];
                alert("the state.user is now " + state.user)
            })
            //createDeposit
            .addCase(createDepositAsync.pending, (state) => {
                state.status = 'loading';
            })
            .addCase(createDepositAsync.fulfilled, (state, action) => {
                state.status = 'idle';
                state.user = action.payload["error"]["name"];
                alert("the state.user is now " + state.user)
            })
            //createLogin
            .addCase(createLoginAsync.pending, (state) => {
                state.status = 'loading';
            })
            .addCase(createLoginAsync.fulfilled, (state, action) => {
                state.status = 'idle';
                state.user = action.payload["response"]["token"];
                state.loggedIn = true;
                alert("the state.user is now " + state.user)
            })
            .addCase(createLoginAsync.rejected, (state, action) => {
                state.status = 'failed';
                alert("login rejected " + action.payload)
                alert("the state.user is now " + state.user)
            })
            //createUser
            .addCase(createUserAsync.pending, (state) => {
                state.status = 'loading';
            })
            .addCase(createUserAsync.fulfilled, (state, action) => {
                state.status = 'idle';
                state.message = action.payload["response"];
                state.loggedIn = true;
                alert("the state.message is now " + state.message)
            })
            .addCase(createUserAsync.rejected, (state, action) => {
                state.status = 'failed';
                alert("createUser rejected " + action.payload)
                alert("the state.message is now " + state.message)
            })
        ;
    },
});

export const {
    increment,
    decrement,
    incrementByAmount,
    setUser,
    setAmount,
    resetState,
    createUser,
    makeTransaction,
    makeLogin,
    makeDeposit,
    makeDelete,
    makeInfo
} = bankingSlice.actions;

// The function below is called a selector and allows us to select a value from
// the state. Selectors can also be defined inline where they're used instead of
// in the slice file. For example: `useSelector((state: RootState) => state.counter.value)`
export const selectBanking = (state: RootState) => state.banking.amount;
export const selectBankingUser = (state: RootState) => state.banking.user;
export const selectLoggedIn = (state: RootState) => state.banking.loggedIn;

// We can also write thunks by hand, which may contain both sync and async logic.
// Here's an example of conditionally dispatching actions based on current state.
export const incrementIfOdd = (amount: number): AppThunk => (
    dispatch,
    getState
) => {
    const currentValue = selectBanking(getState());
    if (currentValue % 2 === 1) {
        dispatch(incrementByAmount(amount));
    }
};

export default bankingSlice.reducer;