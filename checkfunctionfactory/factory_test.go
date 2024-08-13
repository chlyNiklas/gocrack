package checkfunctionfactory

import (
	"testing"
)

func TestNew_IfInvalidHexstring_ErrorNotNil(t *testing.T) {
	// Test testCases
	testCases := []struct {
		hexString string
		isValid   bool
	}{
		{
			"asdfghjkll",
			false,
		},
		{
			"4aba6d9720fcfd366704f40eb3093d168ca4a898",
			true,
		},
		{
			"asdf3ghj1kll",
			false,
		},
		{
			"x",
			false,
		},
		{
			"cb5ff48100d7a36e1aca3ad878392ef752f178ac3b89fa55e88d02374844434d45329f519856987a3fe1ca5589c14164a26e2c508e392d4f3163d15c4d347890",
			true,
		},
		{
			"a234567890'sdfghjkll",
			false,
		},
	}

	for _, testCase := range testCases {
		_, err := New(Md5, testCase.hexString)

		if (err == nil) != testCase.isValid {
			t.Errorf("Test failed for hexString: %s, expected valid: %v", testCase.hexString, testCase.isValid)
		}
	}
}

func TestCreateCheckFungtion_ReturnedFunctionWorks(t *testing.T) {
	testCases := []struct {
		hashString string
		algorithm  HashAlgorithm
		testString string
		maching    bool
	}{
		{
			"038703c7230ae012e3c783ace1d09d64",
			Md5,
			"works",
			true,
		},
		{
			"038703c7230ae012e3c783ace1d09d64",
			Md5,
			"worksnot",
			false,
		},
		{
			"4aba6d9720fcfd366704f40eb3093d168ca4a898",
			Sha1,
			"doesntwork",
			false,
		},
		{
			"295b6b7ea32a096943c54178b4411f93e8a864d6",
			Sha1,
			"works",
			true,
		},
		{
			"93f6e11d9ce3e407e0eef9c87aacabf3f31235a92403845bb5093b4c6f01f2e5",
			Sha256,
			"why doesn't this work?",
			false,
		},
		{
			"ba4c8ceb46a371105969f90c160c1af67535b18e316db0356dc841b89d5b60d9",
			Sha256,
			"maby now?",
			true,
		},
		{
			"1d5468d08363b78b1d2f03565b17b9d317f17d7fa922e589d54073cca37b9aba7ed8e09816dfbdf2f83b02e8553286e12360e4e0b5414c6fb33e471739d28522",
			Sha512,
			"again?",
			false,
		},
		{
			"560296f57c7821d4b3bfeca4bfd9e68d5f8c6bfdb57a6d858ba451ef382be7b92c45ba9d608995164d22cdb984b2762769b6e689e58f32c26c4f24a39f06bc31",
			Sha512,
			"finaly",
			true,
		},
	}

	for _, testCase := range testCases {
		factory, err := New(testCase.algorithm, testCase.hashString)
		checkFunc, err := factory.CreateCheckFunction()
		if err != nil {
			t.Fatal()
		}

		if checkFunc([]byte(testCase.testString)) != testCase.maching {
			t.Errorf(
				"%s is equal to %s when hashed with %s should have been %t, but wasn't",
				testCase.testString,
				testCase.hashString,
				string(testCase.algorithm),
				testCase.maching,
			)
		}

	}

}
