import counterReducer, {
    BankingState,
    increment,
    decrement,
    incrementByAmount,
} from './BankingSlice';

describe('counter reducer', () => {
    const initialState: BankingState = {
        amount: 0,
        user: 'david',
        password: '12345',
        destination: 'allie',
        message: 'msg',
        loggedIn: false,
        status: 'idle'
    };
    it('should handle initial state', () => {
        expect(counterReducer(undefined, { type: 'unknown' })).toEqual({
            value: 0,
            status: 'idle',
        });
    });

    it('should handle increment', () => {
        const actual = counterReducer(initialState, increment());
        expect(actual.amount).toEqual(4);
    });

    it('should handle decrement', () => {
        const actual = counterReducer(initialState, decrement());
        expect(actual.amount).toEqual(2);
    });

    it('should handle incrementByAmount', () => {
        const actual = counterReducer(initialState, incrementByAmount(2));
        expect(actual.amount).toEqual(5);
    });
});
