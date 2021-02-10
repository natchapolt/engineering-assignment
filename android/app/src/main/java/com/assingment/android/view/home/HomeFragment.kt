package com.assingment.android.view.home

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.EditText
import androidx.fragment.app.Fragment
import androidx.fragment.app.activityViewModels
import androidx.navigation.fragment.findNavController
import com.assingment.android.R
import com.assingment.android.view.ShareViewModel

class HomeFragment: Fragment() {

    private val layoutId: Int = R.layout.fragment_home

    private val model: ShareViewModel by activityViewModels()

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(layoutId, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val button = view.findViewById<Button>(R.id.ctaButton)
        val editText = view.findViewById<EditText>(R.id.nameInput)

        button.setOnClickListener {
            model.name.value = editText.text.toString()
            findNavController().navigate(R.id.action_homeFragment_to_quizFragment)
        }
    }
}
