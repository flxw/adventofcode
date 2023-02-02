#!/usr/bin/env python

import solution
import unittest
import copy

class TestSolution(unittest.TestCase):
    INPUT="""    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
"""
    EXPECTED_INITIAL_STATE = [['Z', 'N'], ['M', 'C', 'D'], ['P']]
    EXPECTED_INTERMEDIATE_STATE=[['Z','N','D','C'], ['M'], ['P']]
    EXPECTED_INTERMEDIATE_STATE9001=[['Z','N','C','D'], ['M'], ['P']]
    EXPECTED_INSTRUCTIONS = [
        "move 1 from 2 to 1",
        "move 3 from 1 to 3",
        "move 2 from 2 to 1",
        "move 1 from 1 to 2"
    ]

    def test_parsing(self):
        split_input = TestSolution.INPUT.splitlines()
        initial_state, instructions = solution.parse_input(split_input)

        self.assertEqual(initial_state, TestSolution.EXPECTED_INITIAL_STATE)
        self.assertEqual(instructions, TestSolution.EXPECTED_INSTRUCTIONS)

    def test_execution_of_single_instruction(self):
        initial_state = copy.deepcopy(TestSolution.EXPECTED_INITIAL_STATE)
        intermediate_state = solution.execute_instruction(
            initial_state,
            "move 2 from 2 to 1"
        )
        self.assertEqual(intermediate_state, TestSolution.EXPECTED_INTERMEDIATE_STATE)

    def test_execution_of_single_instruction9001(self):
        initial_state = copy.deepcopy(TestSolution.EXPECTED_INITIAL_STATE)
        intermediate_state = solution.execute_instruction9001(
            initial_state,
            "move 2 from 2 to 1"
        )
        self.assertEqual(intermediate_state, TestSolution.EXPECTED_INTERMEDIATE_STATE9001)

    def test_last_crate_string(self):
        lcs = solution.last_crate_string(TestSolution.EXPECTED_INITIAL_STATE)
        self.assertEqual(lcs, 'NDP')


if __name__ == '__main__':
    unittest.main()