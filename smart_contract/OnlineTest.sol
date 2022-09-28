// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract OnlineTest {

    struct Exam {
        address adminAddr; // address admin 
        uint reward;    // reward
        string answer;  // answer
        uint userSubmitFee;// Admin decide how much fee user should pay to submit their answer
        EXAM_STATE examState; //state of exam


    }
    enum EXAM_STATE { OPEN, CLOSED, ENDSUBMIT, CACULATING_WINNER } // State of the exam
    struct UserExam { 
        address userAddr;  // address of user
        uint examCode;  // code of exam that user participate
        string answer;  // answer
        uint submitTime;    // time submit
    }
    struct Winner {
        address winnerAddr; // Winnner address
        uint submitTime;    // time submit
    }

    address public contractOwner; // Address of contract owner
    address public adminAddr; //admin who create the exam
    uint public adminEntranceFee; // contract owner decide how much fee admin need to pay when create exam
    uint public examFund;   // Fund of the exam
    uint[] public examCodes; // Array of exam code
    uint public examCode=0; //initialize exam code
    Winner[] winners;
    Winner[] clear;
    UserExam[] public userExams;    // array of user Exam
    mapping(uint => Exam) public exams; // Mapping examcode to Exam
    constructor(){
        contractOwner = msg.sender; // Init adress of contract owner
    }
    function chooseAdmin(address _adminAddr,uint _adminEntranceFee)public{
        require(msg.sender==contractOwner, "To choose Admin you must be contract owner!");
        adminAddr=_adminAddr; // init Admin for creating the exam
        adminEntranceFee=_adminEntranceFee;// init admin entranceFee
    }
   
    modifier satisfyCondition() {   // Check satisfy all condition and ready to create Exam
        require(adminAddr !=address(0));  //already choose admin
        require(adminEntranceFee!=0);//already init entranceFee
        _;
    }

    function createExams(uint _reward,uint _userSubmitFee) public payable satisfyCondition{ // Create Exams 
        require(msg.sender==adminAddr,"You must be Admin to Create Exam");
        require(msg.value >= (_reward+adminEntranceFee),"Not enough Ether to Create Exam and Pay Winner");
        examCode++; // generate new examCode
        examCodes.push(examCode);// push examCode to array
        exams[examCode].adminAddr = msg.sender; //update admin address of the exams
        exams[examCode].reward = _reward;  //set reward
        exams[examCode].examState=EXAM_STATE.OPEN; //set state
        exams[examCode].userSubmitFee=_userSubmitFee; // set user submit fee
        examFund += adminEntranceFee;
        payable(address(this)).send(adminEntranceFee); //send admin fee to this contract
        //TODO: maybe tranfer fee to contract here
    }

    function submitTest(uint _examCode,string memory _answer) public payable { // user submit their exam
        require(exams[_examCode].examState==EXAM_STATE.OPEN,"Cannot submit yet!");
        require(msg.value >= exams[_examCode].userSubmitFee* 10**18,"Not enough Ether to submit your exam!");  
        UserExam memory userExam;
        userExam.userAddr = msg.sender;
        userExam.examCode = _examCode;
        userExam.answer = _answer;
        userExam.submitTime = block.timestamp;
        userExams.push(userExam);
        examFund += exams[_examCode].userSubmitFee;
        payable(address(this)).send(exams[_examCode].userSubmitFee); //send user submit fee to this contract
        //TODO: maybe user  need to tranfer fee here
    }

    function endSubmit(uint _examCode) public { // admin end all user submit
        require(msg.sender==exams[_examCode].adminAddr,"You need to be Admin to end submit!");
        require(exams[_examCode].examState==EXAM_STATE.OPEN,"Admin Can't end submit yet!");
        exams[_examCode].examState=EXAM_STATE.ENDSUBMIT; // Change state to ENDSUBMIT
    }

    function submitAnswer(uint _examCode, string memory _answer) public payable { // Admin submit answer
        require(msg.sender==exams[_examCode].adminAddr,"You need to be Admin to submit answer!");
        require(exams[_examCode].examState==EXAM_STATE.ENDSUBMIT,"Admin Can't submit answer yet!");
        exams[_examCode].answer = _answer;
        exams[_examCode].examState=EXAM_STATE.CACULATING_WINNER; // Change state to Caculate winner
    }
    function caculateWinner(uint _examCode) public { // Caculate winner
        require(exams[_examCode].adminAddr == msg.sender,"Only admin can caculate winner!");
        for (uint i = 0; i < userExams.length ; i++) { //travels through array userExams
            if(userExams[i].examCode == _examCode){ //check userExams in which exams
                if((keccak256(abi.encodePacked(userExams[i].answer))) == (keccak256(abi.encodePacked(exams[_examCode].answer)))){
                    // check equal answer
                    Winner memory winner;
                    winner.winnerAddr = userExams[i].userAddr;
                    winner.submitTime = userExams[i].submitTime;
                    winners.push(winner);
                }
            }
        }
        uint tg;
        for (uint i = 0; i < winners.length - 1; i++){
            for (uint j = i + 1; j < winners.length; j ++){
                if(winners[i].submitTime > winners[j].submitTime){  //sort in list winner by time
                    tg = winners[i].submitTime;
                    winners[i].submitTime = winners[j].submitTime;
                    winners[j].submitTime = tg;
                }
            }
        } 
        if(winners.length >= 10) {//mode 1
            uint _reward = exams[_examCode].reward / 10;
            for(uint i = 0; i < 10; i++){
                payable(winners[i].winnerAddr).transfer(_reward);
            }
        }else {//mode 2
            uint _reward = exams[_examCode].reward / winners.length;
            for(uint i = 0; i <winners.length; i++) {
                payable(winners[i].winnerAddr).transfer(_reward);
            }
        }
        //reset
        winners = clear;
        exams[_examCode].examState = EXAM_STATE.CLOSED;
    
    }
    // Function to receive Ether. msg.data must be empty
    receive() external payable {}

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
    function getAccountBalance() public view returns (uint){
        return msg.sender.balance;
    }
     function viewValue(uint _reward,uint _userSubmitFee) public view returns (uint){
        return (_reward+adminEntranceFee)* 10**18;
    }

    function withdraw() public { // with draw fund exam
        require(msg.sender == contractOwner);
        payable(contractOwner).transfer(examFund); 
    }

    
}