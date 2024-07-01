package serializer_test

import (
	"testing"

	"github.com/haiyen11231/pcbook/sample"

	"google.golang.org/protobuf/proto"

	"github.com/haiyen11231/pcbook/serializer"
	"github.com/stretchr/testify/require"
)

func testFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	er := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
