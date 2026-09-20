package store

import (
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

// WeiString renders a NUMERIC(38,0) wei column as a plain decimal string,
// e.g. "25000000000000000000000". All wei amounts in this schema have scale
// 0, so this never needs to handle a fractional part.
func WeiString(n pgtype.Numeric) string {
	if !n.Valid || n.Int == nil {
		return "0"
	}
	b, err := n.MarshalJSON()
	if err != nil {
		return "0"
	}
	return string(b)
}

// WeiFromBigInt converts an on-chain uint256 wei amount into a NUMERIC(38,0)
// value ready to bind as a query parameter.
func WeiFromBigInt(v *big.Int) pgtype.Numeric {
	if v == nil {
		return pgtype.Numeric{Valid: false}
	}
	return pgtype.Numeric{Int: new(big.Int).Set(v), Exp: 0, Valid: true}
}
