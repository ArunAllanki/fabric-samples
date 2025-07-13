const fs = require("fs");
const path = require("path");
const { Wallets } = require("fabric-network");

async function main() {
  const walletPath = path.join(__dirname, "wallet");
  const wallet = await Wallets.newFileSystemWallet(walletPath);

  const certDir = path.resolve(
    __dirname,
    "../test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts"
  );
  const keyDir = path.resolve(
    __dirname,
    "../test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore"
  );

  const certFile = fs.readdirSync(certDir)[0];
  const keyFile = fs.readdirSync(keyDir)[0];

  const cert = fs.readFileSync(path.join(certDir, certFile)).toString();
  const key = fs.readFileSync(path.join(keyDir, keyFile)).toString();

  await wallet.put("admin", {
    credentials: {
      certificate: cert,
      privateKey: key,
    },
    mspId: "Org1MSP",
    type: "X.509",
  });

  console.log("✅ Admin identity imported successfully to wallet");
}

main();
