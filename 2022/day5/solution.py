#!/usr/bin/env python

def parse_input(feed):
    initial_state_description = []
    
    while "[" in feed[0]:
        initial_state_description.append(feed.pop(0))

    slot_indices = feed.pop(0)
    number_of_slots = len([int(idx) for idx in slot_indices.split()])
    feed.pop(0) # pop away the whitespace line

    initial_state = parse_initial_state_description(initial_state_description, number_of_slots)

    return initial_state, feed

def parse_initial_state_description(desc, n_slots):
    state = [list() for x in range(n_slots)]
    characters_per_line = 4 * n_slots - 1

    for line in desc:
        for idx, i in enumerate(range(1, characters_per_line, 4)):
            segment = line[i:i+1]
            if segment.strip():
                state[idx].append(segment)

    for l in state:
        l.reverse()

    return state

def execute_instructions(state, instructions):
    for instruction in instructions:
        execute_instruction(state, instruction)

    return state

def execute_instruction(state, instruction):
    # the structure is always the same
    # move X from Y to Z
    _, amount, _, origin, _, destination = instruction.split()
    amount = int(amount)
    origin = int(origin)
    destination = int(destination)

    for _ in range(amount):
        state[destination-1].append(state[origin-1].pop())

    return state

def execute_instructions9001(state, instructions):
    for instruction in instructions:
        execute_instruction9001(state, instruction)

    return state

def execute_instruction9001(state, instruction):
    # the structure is always the same
    # move X from Y to Z
    _, amount, _, origin, _, destination = instruction.split()
    amount = int(amount)
    origin = int(origin)
    destination = int(destination)

    sublist = []
    for _ in range(amount):
        sublist.append(state[origin-1].pop())

    sublist.reverse()
    state[destination-1].extend(sublist)

    return state


def last_crate_string(state):
    lcs = []

    for stack in state:
        lcs.append(stack[-1])
    
    return ''.join(lcs)


if __name__ == '__main__':
    print("Assuming a crate mover 9000")
    with open('day5/input.txt', 'rt') as file:
        lines = [line for line in file]
        initial_state, instructions = parse_input(lines)
        final_state = execute_instructions(initial_state, instructions)
        lcs = last_crate_string(final_state)
        print(lcs)

    print("Assuming a crate mover 9001")
    with open('day5/input.txt', 'rt') as file:
        lines = [line for line in file]
        initial_state, instructions = parse_input(lines)
        final_state = execute_instructions9001(initial_state, instructions)
        lcs = last_crate_string(final_state)
        print(lcs)