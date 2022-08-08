package serializer_test

import (
	"testing"

	"example.com/grpc/pb"
	"example.com/grpc/sample"
	"example.com/grpc/serializer"
	proto "github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop := sample.NewLaptop()

	err := serializer.WriteProtobufToBinary(laptop, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop, jsonFile)
	require.NoError(t, err)

	laptop1 := &pb.Laptop{}

	err = serializer.ReadProtobuffFromBinay(laptop1, binaryFile)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop, laptop1))
}
