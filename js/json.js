(function(){
	var stack = [];
	var curObj = null;
	var curValue=null;

	curObj={};
	stack.push(curObj);
	curObj=[];
	stack.push(curObj);
	curValue=1
	curObj.push(curValue);
	curValue=2
	curObj.push(curValue);
	curValue=3
	curObj.push(curValue);
	curObj=stack.pop();
	curValue=curObj;
	curObj=stack[stack.length-1];

	curObj["a"]=curValue;
	curObj=[];
	stack.push(curObj);
	curValue="a\tbc"
	curObj.push(curValue);
	curValue="12  3"
	curObj.push(curValue);
	curValue="4,5\"6"
	curObj.push(curValue);
	curObj={};
	stack.push(curObj);
	curValue=1
	curObj["x"]=curValue;
	curValue="cc\ncc"
	curObj["y"]=curValue;
	curObj=stack.pop();
	curValue=curObj;
	curObj=stack[stack.length-1];

	curObj.push(curValue);
	curValue=4.56
	curObj.push(curValue);
	curObj=stack.pop();
	curValue=curObj;
	curObj=stack[stack.length-1];

	curObj["b"]=curValue;
	curValue="I'm OK~"
	curObj["text"]=curValue;
	curValue=234
	curObj["1-2"]=curValue;
	curObj=stack.pop();
	curValue=curObj;
	curObj=stack[stack.length-1];

	return curValue;
})()