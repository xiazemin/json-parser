package antlt;
// Generated from /Users/didi/PhpstormProjects/c/json-parser/antlr/calc/src/main/java/AntlrTest.g4 by ANTLR 4.8
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link AntlrTestParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface AntlrTestVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link AntlrTestParser#prog}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitProg(AntlrTestParser.ProgContext ctx);
	/**
	 * Visit a parse tree produced by the {@code print}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrint(AntlrTestParser.PrintContext ctx);
	/**
	 * Visit a parse tree produced by the {@code assign}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssign(AntlrTestParser.AssignContext ctx);
	/**
	 * Visit a parse tree produced by the {@code blank}
	 * labeled alternative in {@link AntlrTestParser#stat}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBlank(AntlrTestParser.BlankContext ctx);
	/**
	 * Visit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMulDiv(AntlrTestParser.MulDivContext ctx);
	/**
	 * Visit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAddSub(AntlrTestParser.AddSubContext ctx);
	/**
	 * Visit a parse tree produced by the {@code IdInt}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdInt(AntlrTestParser.IdIntContext ctx);
	/**
	 * Visit a parse tree produced by the {@code Private}
	 * labeled alternative in {@link AntlrTestParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrivate(AntlrTestParser.PrivateContext ctx);
	/**
	 * Visit a parse tree produced by {@link AntlrTestParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValue(AntlrTestParser.ValueContext ctx);
}