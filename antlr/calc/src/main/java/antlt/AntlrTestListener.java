package antlt;
// Generated from /Users/didi/PhpstormProjects/c/json-parser/antlr/calc/src/main/java/AntlrTest.g4 by ANTLR 4.8
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link AntlrTestParser}.
 */
public interface AntlrTestListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link AntlrTestParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(AntlrTestParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link AntlrTestParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(AntlrTestParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by the {@code print}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 */
	void enterPrint(AntlrTestParser.PrintContext ctx);
	/**
	 * Exit a parse tree produced by the {@code print}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 */
	void exitPrint(AntlrTestParser.PrintContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assign}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 */
	void enterAssign(AntlrTestParser.AssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assign}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 */
	void exitAssign(AntlrTestParser.AssignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code blank}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 */
	void enterBlank(AntlrTestParser.BlankContext ctx);
	/**
	 * Exit a parse tree produced by the {@code blank}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 */
	void exitBlank(AntlrTestParser.BlankContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterMulDiv(AntlrTestParser.MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitMulDiv(AntlrTestParser.MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(AntlrTestParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(AntlrTestParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdInt}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIdInt(AntlrTestParser.IdIntContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdInt}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIdInt(AntlrTestParser.IdIntContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Private}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterPrivate(AntlrTestParser.PrivateContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Private}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitPrivate(AntlrTestParser.PrivateContext ctx);
	/**
	 * Enter a parse tree produced by {@link AntlrTestParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(AntlrTestParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link AntlrTestParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(AntlrTestParser.ValueContext ctx);
}