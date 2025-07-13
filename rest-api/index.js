const express = require("express");
const cors = require("cors");
const fs = require("fs");
const path = require("path");
const { Gateway, Wallets } = require("fabric-network");

const app = express();
app.use(cors());
app.use(express.json());

const ccpPath = path.resolve(__dirname, "connection.json");
const ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));

async function getContract() {
  const wallet = await Wallets.newFileSystemWallet(path.join(__dirname, "wallet"));
  const gateway = new Gateway();
  await gateway.connect(ccp, {
    wallet,
    identity: "admin",
    discovery: { enabled: true, asLocalhost: true },
  });
  const network = await gateway.getNetwork("mychannel");
  const contract = network.getContract("accounts");
  return { gateway, contract };
}

app.post("/account", async (req, res) => {
  try {
    const {
      dealerId, msisdn, mpin, balance,
      status, transAmount, transType, remarks
    } = req.body;

    const { gateway, contract } = await getContract();
    await contract.submitTransaction("CreateAccount",
      dealerId, msisdn, mpin,
      balance.toString(), status,
      transAmount.toString(), transType, remarks);
    await gateway.disconnect();
    res.status(200).json({ message: "Account created" });
  } catch (e) {
    res.status(500).json({ error: e.message });
  }
});

app.get("/account/:id", async (req, res) => {
  try {
    const { gateway, contract } = await getContract();
    const result = await contract.evaluateTransaction("ReadAccount", req.params.id);
    await gateway.disconnect();
    res.status(200).json(JSON.parse(result.toString()));
  } catch (e) {
    res.status(404).json({ error: e.message });
  }
});

app.listen(4000, () => {
  console.log("REST API running at http://localhost:4000");
});
