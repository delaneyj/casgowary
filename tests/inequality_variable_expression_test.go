package tests

// public class VariableExpression {
//     private static double EPSILON = 1.0e-8;

//     @Test
//     public void lessThanEqualTo() throws DuplicateConstraintException, UnsatisfiableConstraintException {
//         Variable x = new Variable("x");
//         Solver solver = new Solver();
//         solver.addConstraint(Symbolics.lessThanOrEqualTo(x, new Expression(100)));
//         solver.updateVariables();
//         assertTrue(x.getValue() <= 100);
//         solver.addConstraint(Symbolics.equals(x, 90));
//         solver.updateVariables();
//         assertEquals(x.getValue(), 90, EPSILON);
//     }

//     @Test(expected = UnsatisfiableConstraintException.class)
//     public void lessThanEqualToUnsatisfiable() throws DuplicateConstraintException, UnsatisfiableConstraintException {
//         Variable x = new Variable("x");
//         Solver solver = new Solver();
//         solver.addConstraint(Symbolics.lessThanOrEqualTo(x, new Expression(100)));
//         solver.updateVariables();
//         assertTrue(x.getValue() <= 100);
//         solver.addConstraint(Symbolics.equals(x, 110));
//         solver.updateVariables();
//     }

//     @Test
//     public void greaterThanEqualTo() throws DuplicateConstraintException, UnsatisfiableConstraintException {
//         Variable x = new Variable("x");
//         Solver solver = new Solver();
//         solver.addConstraint(Symbolics.greaterThanOrEqualTo(x, new Expression(100)));
//         solver.updateVariables();
//         assertTrue(x.getValue() >= 100);
//         solver.addConstraint(Symbolics.equals(x, 110));
//         solver.updateVariables();
//         assertEquals(x.getValue(), 110, EPSILON);
//     }

//     @Test(expected = UnsatisfiableConstraintException.class)
//     public void greaterThanEqualToUnsatisfiable() throws DuplicateConstraintException, UnsatisfiableConstraintException {
//         Variable x = new Variable("x");
//         Solver solver = new Solver();
//         solver.addConstraint(Symbolics.greaterThanOrEqualTo(x, 100));
//         solver.updateVariables();
//         assertTrue(x.getValue() >= 100);
//         solver.addConstraint(Symbolics.equals(x, 90));
//         solver.updateVariables();
//     }
// }
