package antlt;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CodePointCharStream;
import org.antlr.v4.runtime.CommonTokenStream;

public class Main {

    public static void run(String expr) throws Exception {

        //对每一个输入的字符串，构造一个 CodePointCharStream
        CodePointCharStream cpcs = CharStreams.fromString(expr);

        //用 cpcs 构造词法分析器 lexer，词法分析的作用是产生记号
        AntlrTestLexer lexer = new AntlrTestLexer(cpcs);

        //用词法分析器 lexer 构造一个记号流 tokens
        CommonTokenStream tokens = new CommonTokenStream(lexer);

        //再使用 tokens 构造语法分析器 parser,至此已经完成词法分析和语法分析的准备工作
        AntlrTestParser parser = new AntlrTestParser(tokens);

        //最终调用语法分析器的规则 prog，完成对表达式的验证
        AntlrTestParser.ProgContext pcontext = parser.prog();

        // 通过访问者模式，执行我们的程序
        EvalVisitor evalVisitor = new EvalVisitor();
        evalVisitor.visit(pcontext);
    }

    public static void main(String[] args) throws Exception {
        run("a=5\nb=3\nc=a*b+3\nc*c\n");
    }
}