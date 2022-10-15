// Code generated by fastssz. DO NOT EDIT.
// Hash: 960775919c6966b4ccb2db33ac2493ebb862360f0a433c96a256085b8f83198d
// Version: 0.1.2
package cltypes

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the MetadataV1 object
func (m *MetadataV1) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MetadataV1 object to a target array
func (m *MetadataV1) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SeqNumber'
	dst = ssz.MarshalUint64(dst, m.SeqNumber)

	// Field (1) 'Attnets'
	dst = ssz.MarshalUint64(dst, m.Attnets)

	return
}

// UnmarshalSSZ ssz unmarshals the MetadataV1 object
func (m *MetadataV1) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'SeqNumber'
	m.SeqNumber = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Attnets'
	m.Attnets = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MetadataV1 object
func (m *MetadataV1) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the MetadataV1 object
func (m *MetadataV1) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MetadataV1 object with a hasher
func (m *MetadataV1) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SeqNumber'
	hh.PutUint64(m.SeqNumber)

	// Field (1) 'Attnets'
	hh.PutUint64(m.Attnets)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the MetadataV1 object
func (m *MetadataV1) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}

// MarshalSSZ ssz marshals the MetadataV2 object
func (m *MetadataV2) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MetadataV2 object to a target array
func (m *MetadataV2) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SeqNumber'
	dst = ssz.MarshalUint64(dst, m.SeqNumber)

	// Field (1) 'Attnets'
	dst = ssz.MarshalUint64(dst, m.Attnets)

	// Field (2) 'Syncnets'
	dst = ssz.MarshalUint64(dst, m.Syncnets)

	return
}

// UnmarshalSSZ ssz unmarshals the MetadataV2 object
func (m *MetadataV2) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 24 {
		return ssz.ErrSize
	}

	// Field (0) 'SeqNumber'
	m.SeqNumber = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Attnets'
	m.Attnets = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'Syncnets'
	m.Syncnets = ssz.UnmarshallUint64(buf[16:24])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MetadataV2 object
func (m *MetadataV2) SizeSSZ() (size int) {
	size = 24
	return
}

// HashTreeRoot ssz hashes the MetadataV2 object
func (m *MetadataV2) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MetadataV2 object with a hasher
func (m *MetadataV2) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SeqNumber'
	hh.PutUint64(m.SeqNumber)

	// Field (1) 'Attnets'
	hh.PutUint64(m.Attnets)

	// Field (2) 'Syncnets'
	hh.PutUint64(m.Syncnets)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the MetadataV2 object
func (m *MetadataV2) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}

// MarshalSSZ ssz marshals the ENRForkID object
func (e *ENRForkID) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ENRForkID object to a target array
func (e *ENRForkID) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentForkDigest'
	dst = append(dst, e.CurrentForkDigest[:]...)

	// Field (1) 'NextForkVersion'
	dst = append(dst, e.NextForkVersion[:]...)

	// Field (2) 'NextForkEpoch'
	dst = ssz.MarshalUint64(dst, e.NextForkEpoch)

	return
}

// UnmarshalSSZ ssz unmarshals the ENRForkID object
func (e *ENRForkID) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentForkDigest'
	copy(e.CurrentForkDigest[:], buf[0:4])

	// Field (1) 'NextForkVersion'
	copy(e.NextForkVersion[:], buf[4:8])

	// Field (2) 'NextForkEpoch'
	e.NextForkEpoch = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ENRForkID object
func (e *ENRForkID) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the ENRForkID object
func (e *ENRForkID) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ENRForkID object with a hasher
func (e *ENRForkID) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentForkDigest'
	hh.PutBytes(e.CurrentForkDigest[:])

	// Field (1) 'NextForkVersion'
	hh.PutBytes(e.NextForkVersion[:])

	// Field (2) 'NextForkEpoch'
	hh.PutUint64(e.NextForkEpoch)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ENRForkID object
func (e *ENRForkID) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}

// MarshalSSZ ssz marshals the ForkData object
func (f *ForkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the ForkData object to a target array
func (f *ForkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentVersion'
	dst = append(dst, f.CurrentVersion[:]...)

	// Field (1) 'GenesisValidatorsRoot'
	dst = append(dst, f.GenesisValidatorsRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ForkData object
func (f *ForkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 36 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[0:4])

	// Field (1) 'GenesisValidatorsRoot'
	copy(f.GenesisValidatorsRoot[:], buf[4:36])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ForkData object
func (f *ForkData) SizeSSZ() (size int) {
	size = 36
	return
}

// HashTreeRoot ssz hashes the ForkData object
func (f *ForkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the ForkData object with a hasher
func (f *ForkData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	hh.PutBytes(f.CurrentVersion[:])

	// Field (1) 'GenesisValidatorsRoot'
	hh.PutBytes(f.GenesisValidatorsRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ForkData object
func (f *ForkData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(f)
}

// MarshalSSZ ssz marshals the Ping object
func (p *Ping) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the Ping object to a target array
func (p *Ping) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Id'
	dst = ssz.MarshalUint64(dst, p.Id)

	return
}

// UnmarshalSSZ ssz unmarshals the Ping object
func (p *Ping) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 8 {
		return ssz.ErrSize
	}

	// Field (0) 'Id'
	p.Id = ssz.UnmarshallUint64(buf[0:8])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Ping object
func (p *Ping) SizeSSZ() (size int) {
	size = 8
	return
}

// HashTreeRoot ssz hashes the Ping object
func (p *Ping) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the Ping object with a hasher
func (p *Ping) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Id'
	hh.PutUint64(p.Id)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Ping object
func (p *Ping) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(p)
}

// MarshalSSZ ssz marshals the SingleRoot object
func (s *SingleRoot) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SingleRoot object to a target array
func (s *SingleRoot) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Root'
	dst = append(dst, s.Root[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SingleRoot object
func (s *SingleRoot) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 32 {
		return ssz.ErrSize
	}

	// Field (0) 'Root'
	copy(s.Root[:], buf[0:32])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SingleRoot object
func (s *SingleRoot) SizeSSZ() (size int) {
	size = 32
	return
}

// HashTreeRoot ssz hashes the SingleRoot object
func (s *SingleRoot) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SingleRoot object with a hasher
func (s *SingleRoot) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Root'
	hh.PutBytes(s.Root[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SingleRoot object
func (s *SingleRoot) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the LightClientUpdatesByRangeRequest object
func (l *LightClientUpdatesByRangeRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(l)
}

// MarshalSSZTo ssz marshals the LightClientUpdatesByRangeRequest object to a target array
func (l *LightClientUpdatesByRangeRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Period'
	dst = ssz.MarshalUint64(dst, l.Period)

	// Field (1) 'Count'
	dst = ssz.MarshalUint64(dst, l.Count)

	return
}

// UnmarshalSSZ ssz unmarshals the LightClientUpdatesByRangeRequest object
func (l *LightClientUpdatesByRangeRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'Period'
	l.Period = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Count'
	l.Count = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the LightClientUpdatesByRangeRequest object
func (l *LightClientUpdatesByRangeRequest) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the LightClientUpdatesByRangeRequest object
func (l *LightClientUpdatesByRangeRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(l)
}

// HashTreeRootWith ssz hashes the LightClientUpdatesByRangeRequest object with a hasher
func (l *LightClientUpdatesByRangeRequest) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Period'
	hh.PutUint64(l.Period)

	// Field (1) 'Count'
	hh.PutUint64(l.Count)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the LightClientUpdatesByRangeRequest object
func (l *LightClientUpdatesByRangeRequest) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(l)
}

// MarshalSSZ ssz marshals the Status object
func (s *Status) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the Status object to a target array
func (s *Status) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ForkDigest'
	dst = append(dst, s.ForkDigest[:]...)

	// Field (1) 'FinalizedRoot'
	dst = append(dst, s.FinalizedRoot[:]...)

	// Field (2) 'FinalizedEpoch'
	dst = ssz.MarshalUint64(dst, s.FinalizedEpoch)

	// Field (3) 'HeadRoot'
	dst = append(dst, s.HeadRoot[:]...)

	// Field (4) 'HeadSlot'
	dst = ssz.MarshalUint64(dst, s.HeadSlot)

	return
}

// UnmarshalSSZ ssz unmarshals the Status object
func (s *Status) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 84 {
		return ssz.ErrSize
	}

	// Field (0) 'ForkDigest'
	copy(s.ForkDigest[:], buf[0:4])

	// Field (1) 'FinalizedRoot'
	copy(s.FinalizedRoot[:], buf[4:36])

	// Field (2) 'FinalizedEpoch'
	s.FinalizedEpoch = ssz.UnmarshallUint64(buf[36:44])

	// Field (3) 'HeadRoot'
	copy(s.HeadRoot[:], buf[44:76])

	// Field (4) 'HeadSlot'
	s.HeadSlot = ssz.UnmarshallUint64(buf[76:84])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Status object
func (s *Status) SizeSSZ() (size int) {
	size = 84
	return
}

// HashTreeRoot ssz hashes the Status object
func (s *Status) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the Status object with a hasher
func (s *Status) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ForkDigest'
	hh.PutBytes(s.ForkDigest[:])

	// Field (1) 'FinalizedRoot'
	hh.PutBytes(s.FinalizedRoot[:])

	// Field (2) 'FinalizedEpoch'
	hh.PutUint64(s.FinalizedEpoch)

	// Field (3) 'HeadRoot'
	hh.PutBytes(s.HeadRoot[:])

	// Field (4) 'HeadSlot'
	hh.PutUint64(s.HeadSlot)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Status object
func (s *Status) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the SigningData object
func (s *SigningData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningData object to a target array
func (s *SigningData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Root'
	dst = append(dst, s.Root[:]...)

	// Field (1) 'Domain'
	if size := len(s.Domain); size != 32 {
		err = ssz.ErrBytesLengthFn("SigningData.Domain", size, 32)
		return
	}
	dst = append(dst, s.Domain...)

	return
}

// UnmarshalSSZ ssz unmarshals the SigningData object
func (s *SigningData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'Root'
	copy(s.Root[:], buf[0:32])

	// Field (1) 'Domain'
	if cap(s.Domain) == 0 {
		s.Domain = make([]byte, 0, len(buf[32:64]))
	}
	s.Domain = append(s.Domain, buf[32:64]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SigningData object
func (s *SigningData) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SigningData object
func (s *SigningData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SigningData object with a hasher
func (s *SigningData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Root'
	hh.PutBytes(s.Root[:])

	// Field (1) 'Domain'
	if size := len(s.Domain); size != 32 {
		err = ssz.ErrBytesLengthFn("SigningData.Domain", size, 32)
		return
	}
	hh.PutBytes(s.Domain)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SigningData object
func (s *SigningData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}