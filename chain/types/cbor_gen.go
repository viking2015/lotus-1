package types

import (
	"fmt"
	"io"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

/* This file was generated by github.com/whyrusleeping/cbor-gen */

var _ = xerrors.Errorf

func (t *BlockHeader) MarshalCBOR(w io.Writer) error {
	if _, err := w.Write([]byte{140}); err != nil {
		return err
	}

	// t.t.Miner (address.Address)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.Tickets ([]*types.Ticket)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Tickets)))); err != nil {
		return err
	}
	for _, v := range t.Tickets {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.t.ElectionProof ([]uint8)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.ElectionProof)))); err != nil {
		return err
	}
	if _, err := w.Write(t.ElectionProof); err != nil {
		return err
	}

	// t.t.Parents ([]cid.Cid)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Parents)))); err != nil {
		return err
	}
	for _, v := range t.Parents {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Parents: %w", err)
		}
	}

	// t.t.ParentWeight (types.BigInt)
	if err := t.ParentWeight.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.Height (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, t.Height)); err != nil {
		return err
	}

	// t.t.StateRoot (cid.Cid)
	if err := cbg.WriteCid(w, t.StateRoot); err != nil {
		return xerrors.Errorf("failed to write cid field t.StateRoot: %w", err)
	}

	// t.t.Messages (cid.Cid)
	if err := cbg.WriteCid(w, t.Messages); err != nil {
		return xerrors.Errorf("failed to write cid field t.Messages: %w", err)
	}

	// t.t.BLSAggregate (types.Signature)
	if err := t.BLSAggregate.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.MessageReceipts (cid.Cid)
	if err := cbg.WriteCid(w, t.MessageReceipts); err != nil {
		return xerrors.Errorf("failed to write cid field t.MessageReceipts: %w", err)
	}

	// t.t.Timestamp (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, t.Timestamp)); err != nil {
		return err
	}

	// t.t.BlockSig (types.Signature)
	if err := t.BlockSig.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *BlockHeader) UnmarshalCBOR(br io.Reader) error {

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 12 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Miner (address.Address)

	if err := t.Miner.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.Tickets ([]*types.Ticket)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if extra > 8192 {
		return fmt.Errorf("array too large")
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Tickets = make([]*Ticket, extra)
	}
	for i := 0; i < int(extra); i++ {
		var v Ticket
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Tickets[i] = &v
	}

	// t.t.ElectionProof ([]uint8)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if extra > 8192 {
		return fmt.Errorf("array too large")
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.ElectionProof = make([]byte, extra)
	if _, err := io.ReadFull(br, t.ElectionProof); err != nil {
		return err
	}
	// t.t.Parents ([]cid.Cid)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if extra > 8192 {
		return fmt.Errorf("array too large")
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Parents = make([]cid.Cid, extra)
	}
	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Parents failed: %w", err)
		}
		t.Parents[i] = c
	}

	// t.t.ParentWeight (types.BigInt)

	if err := t.ParentWeight.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.Height (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Height = extra
	// t.t.StateRoot (cid.Cid)

	{
		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.StateRoot: %w", err)
		}
		t.StateRoot = c
	}
	// t.t.Messages (cid.Cid)

	{
		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Messages: %w", err)
		}
		t.Messages = c
	}
	// t.t.BLSAggregate (types.Signature)

	if err := t.BLSAggregate.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.MessageReceipts (cid.Cid)

	{
		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.MessageReceipts: %w", err)
		}
		t.MessageReceipts = c
	}
	// t.t.Timestamp (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Timestamp = extra
	// t.t.BlockSig (types.Signature)

	if err := t.BlockSig.UnmarshalCBOR(br); err != nil {
		return err
	}
	return nil
}

func (t *Ticket) MarshalCBOR(w io.Writer) error {
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.t.VRFProof ([]uint8)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.VRFProof)))); err != nil {
		return err
	}
	if _, err := w.Write(t.VRFProof); err != nil {
		return err
	}

	// t.t.VDFResult ([]uint8)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.VDFResult)))); err != nil {
		return err
	}
	if _, err := w.Write(t.VDFResult); err != nil {
		return err
	}

	// t.t.VDFProof ([]uint8)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.VDFProof)))); err != nil {
		return err
	}
	if _, err := w.Write(t.VDFProof); err != nil {
		return err
	}
	return nil
}

func (t *Ticket) UnmarshalCBOR(br io.Reader) error {

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.VRFProof ([]uint8)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if extra > 8192 {
		return fmt.Errorf("array too large")
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.VRFProof = make([]byte, extra)
	if _, err := io.ReadFull(br, t.VRFProof); err != nil {
		return err
	}
	// t.t.VDFResult ([]uint8)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if extra > 8192 {
		return fmt.Errorf("array too large")
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.VDFResult = make([]byte, extra)
	if _, err := io.ReadFull(br, t.VDFResult); err != nil {
		return err
	}
	// t.t.VDFProof ([]uint8)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if extra > 8192 {
		return fmt.Errorf("array too large")
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.VDFProof = make([]byte, extra)
	if _, err := io.ReadFull(br, t.VDFProof); err != nil {
		return err
	}
	return nil
}

func (t *Message) MarshalCBOR(w io.Writer) error {
	if _, err := w.Write([]byte{136}); err != nil {
		return err
	}

	// t.t.To (address.Address)
	if err := t.To.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.From (address.Address)
	if err := t.From.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.Nonce (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, t.Nonce)); err != nil {
		return err
	}

	// t.t.Value (types.BigInt)
	if err := t.Value.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.GasPrice (types.BigInt)
	if err := t.GasPrice.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.GasLimit (types.BigInt)
	if err := t.GasLimit.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.Method (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, t.Method)); err != nil {
		return err
	}

	// t.t.Params ([]uint8)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Params)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Params); err != nil {
		return err
	}
	return nil
}

func (t *Message) UnmarshalCBOR(br io.Reader) error {

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 8 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.To (address.Address)

	if err := t.To.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.From (address.Address)

	if err := t.From.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.Nonce (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Nonce = extra
	// t.t.Value (types.BigInt)

	if err := t.Value.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.GasPrice (types.BigInt)

	if err := t.GasPrice.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.GasLimit (types.BigInt)

	if err := t.GasLimit.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.Method (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Method = extra
	// t.t.Params ([]uint8)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if extra > 8192 {
		return fmt.Errorf("array too large")
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Params = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Params); err != nil {
		return err
	}
	return nil
}

func (t *SignedMessage) MarshalCBOR(w io.Writer) error {
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.t.Message (types.Message)
	if err := t.Message.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.Signature (types.Signature)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedMessage) UnmarshalCBOR(br io.Reader) error {

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Message (types.Message)

	if err := t.Message.UnmarshalCBOR(br); err != nil {
		return err
	}
	// t.t.Signature (types.Signature)

	if err := t.Signature.UnmarshalCBOR(br); err != nil {
		return err
	}
	return nil
}

func (t *MsgMeta) MarshalCBOR(w io.Writer) error {
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.t.BlsMessages (cid.Cid)
	if err := cbg.WriteCid(w, t.BlsMessages); err != nil {
		return xerrors.Errorf("failed to write cid field t.BlsMessages: %w", err)
	}

	// t.t.SecpkMessages (cid.Cid)
	if err := cbg.WriteCid(w, t.SecpkMessages); err != nil {
		return xerrors.Errorf("failed to write cid field t.SecpkMessages: %w", err)
	}
	return nil
}

func (t *MsgMeta) UnmarshalCBOR(br io.Reader) error {

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.BlsMessages (cid.Cid)

	{
		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.BlsMessages: %w", err)
		}
		t.BlsMessages = c
	}
	// t.t.SecpkMessages (cid.Cid)

	{
		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.SecpkMessages: %w", err)
		}
		t.SecpkMessages = c
	}
	return nil
}